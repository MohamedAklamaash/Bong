To do:

need to infer from a pcap and write a pcap file
    
create a proxy server for only allowing restriced access

create a docker-compose or k8s file to concurrently simulate attack using ddos
    
make a hoic syn flood simulation attack

capture packets that are arriving at a router , get the name of the sender device or the prev hop router name and their corresponding ip

since our sole focus here is to make a anamoly based detection,we arent' using any ml here

Approach that i am going to take

naive approach 1:

    get the respective ip of the senders' device or the prev hop ip
    create a rate limiting window
    if we exceed the rate limiting window
    we will send a mail to the concerned server's authority
    and flag that specific ip by not sending any response from that ip

naive approach 2:

    suppose attackers is maksing all the ip,
    we can pre determine the amount of traffic expected
    we can set an upper bound to process the reqs

    if we somehow reach above the upper bound

    we will process the reqs in such a way that we dont' overhelm the resources of our system
    (i.e..) we will delay the response by not going over the threshold

Any other suggestions are highly welcome