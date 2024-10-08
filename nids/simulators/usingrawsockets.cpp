#include <iostream>
#include <map>
#include <cstring>
#include <arpa/inet.h>
#include <sys/socket.h>
#include <netinet/ip.h>
#include <netinet/tcp.h>
#include <netinet/if_ether.h>  // For Ethernet header
#include <linux/if_ether.h>  
#include <unistd.h>
#include <ctime>

// Thresholds
#define SYN_THRESHOLD 10000
#define DOS_THRESHOLD 30000
#define HTTPS_PORT 443

// IP and SYN counters
std::map<std::string, int> ip_count;
std::map<std::string, int> syn_count;
std::time_t last_check = std::time(0);

// Parse IP address from packet
std::string parse_ip_address(struct iphdr* ip_header) {
    struct in_addr ip_addr;
    ip_addr.s_addr = ip_header->saddr;
    return std::string(inet_ntoa(ip_addr));
}

void analyze_tcp_packet(struct iphdr* ip_header, struct tcphdr* tcp_header) {
    std::string source_ip = parse_ip_address(ip_header);

    // Increment IP packet count
    ip_count[source_ip]++;

    // Detect SYN Flood 
    if (tcp_header->syn && !tcp_header->ack) {
        syn_count[source_ip]++;
    }

    // Detect HTTPS attacks (many connections to port 443)
    if (ntohs(tcp_header->dest) == HTTPS_PORT) {
        std::cout << "Alert: Potential HTTPS attack from " << source_ip << std::endl;
    }

    // Check for SYN flood threshold
    if (syn_count[source_ip] > SYN_THRESHOLD) {
        std::cout << "Alert: Potential SYN flood attack from " << source_ip << std::endl;
        syn_count[source_ip] = 0; // Reset after detection
    }

    // Check for DoS/DDoS attacks 
    if (ip_count[source_ip] > DOS_THRESHOLD) {
        std::cout << "Alert: Potential DoS/DDoS attack from " << source_ip << std::endl;
        ip_count[source_ip] = 0; 
    }
}

void process_packet(char* buffer, int data_size) {
    struct iphdr* ip_header = (struct iphdr*)(buffer + 14);  
    struct tcphdr* tcp_header = (struct tcphdr*)(buffer + ip_header->ihl * 4 + 14); 

    if (ip_header->protocol == IPPROTO_TCP) {
        analyze_tcp_packet(ip_header, tcp_header);
    }
}

// Capturing packets using raw sockets
int main() {
    int raw_socket;
    struct sockaddr saddr;
    socklen_t saddr_len = sizeof(saddr);

    char buffer[65536];  

    raw_socket = socket(AF_PACKET, SOCK_RAW, htons(ETH_P_ALL));
    if (raw_socket < 0) {
        std::cerr << "Error creating socket. Run with root privileges!" << std::endl;
        return 1;
    }

    // Capturing packets indefinitely
    while (true) {
        int data_size = recvfrom(raw_socket, buffer, sizeof(buffer), 0, &saddr, &saddr_len);
        if (data_size < 0) {
            std::cerr << "Error receiving packet." << std::endl;
            close(raw_socket);
            return 1;
        }

        process_packet(buffer, data_size);

        if (std::difftime(std::time(0), last_check) > 10) {
            ip_count.clear();
            syn_count.clear();
            last_check = std::time(0);
        }
    }

    close(raw_socket);
    return 0;
}
