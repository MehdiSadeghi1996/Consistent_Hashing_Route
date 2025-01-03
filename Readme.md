
# Consistent Hashing Route (application level)

an implementation of consistent hashing route with static 6 queue and 6 consumer that guarantee each consumer produce all action on specific entity by id .

for example an order with id = 1 , with 4 action like x,y,w,z  goes to one specific queue and sequentially process by one consumer.





## Diagnostics info

```bash
  docker logs producer
```

2025/01/03 13:22:51 Message sent: &{Id:9991 IsPaid:false Type:X} to queue: ordered_queue_2

2025/01/03 13:22:51 Message sent: &{Id:9991 IsPaid:false Type:Y} to queue: ordered_queue_2

2025/01/03 13:22:51 Message sent: &{Id:9991 IsPaid:false Type:W} to queue: ordered_queue_2

2025/01/03 13:22:51 Message sent: &{Id:9991 IsPaid:false Type:Z} to queue: ordered_queue_2

2025/01/03 13:22:51 Message sent: &{Id:9992 IsPaid:false Type:X} to queue: ordered_queue_3

2025/01/03 13:22:51 Message sent: &{Id:9992 IsPaid:false Type:Y} to queue: ordered_queue_3

2025/01/03 13:22:51 Message sent: &{Id:9992 IsPaid:false Type:W} to queue: ordered_queue_3

2025/01/03 13:22:51 Message sent: &{Id:9992 IsPaid:false Type:Z} to queue: ordered_queue_3

2025/01/03 13:22:51 Message sent: &{Id:9993 IsPaid:false Type:X} to queue: ordered_queue_4

2025/01/03 13:22:51 Message sent: &{Id:9993 IsPaid:false Type:Y} to queue: ordered_queue_4

2025/01/03 13:22:51 Message sent: &{Id:9993 IsPaid:false Type:W} to queue: ordered_queue_4

2025/01/03 13:22:51 Message sent: &{Id:9993 IsPaid:false Type:Z} to queue: ordered_queue_4

2025/01/03 13:22:51 All messages have been sent.

2025/01/03 13:22:51 Shutting down producer...

​

```bash
  docker logs consumer

```


2025/01/03 13:22:51 Processing order: {Id:9962 IsPaid:false Type:X}

2025/01/03 13:22:51 Processing order: {Id:9962 IsPaid:false Type:Y}

2025/01/03 13:22:51 Processing order: {Id:9962 IsPaid:false Type:W}

2025/01/03 13:22:51 Processing order: {Id:9962 IsPaid:false Type:Z}

2025/01/03 13:22:51 Processing order: {Id:9968 IsPaid:false Type:X}

2025/01/03 13:22:51 Processing order: {Id:9968 IsPaid:false Type:Y}

2025/01/03 13:22:51 Processing order: {Id:9968 IsPaid:false Type:W}

2025/01/03 13:22:51 Processing order: {Id:9968 IsPaid:false Type:Z}

2025/01/03 13:22:51 Processing order: {Id:9974 IsPaid:false Type:X}

2025/01/03 13:22:51 Processing order: {Id:9974 IsPaid:false Type:Y}

2025/01/03 13:22:51 Processing order: {Id:9974 IsPaid:false Type:W}

2025/01/03 13:22:51 Processing order: {Id:9974 IsPaid:false Type:Z}
## Deployment

To deploy this project run

```bash
  docker compose up -d --build
```

