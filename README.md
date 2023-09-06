# Go-Chat

GoChat is a micro-blogging platform that allows users to connect with others and engage in real-time chat conversations. The application is built using microservices architecture and offers several key features:

- # Authentication Service:

Users can securely sign up, log in, and log out of the platform.
User registration ensures a secure and personalized experience.
- # User Service:

Users can follow or unfollow other users, fostering social connections within the platform.
Users have the option to delete their account if desired.
- # Real-time Chat:

Users can engage in real-time chat conversations using WebSockets.
Real-time communication enhances user interaction and fosters a dynamic platform experience.
GoChat provides a seamless and interactive chat experience, allowing users to connect and communicate with others in real-time.

***

## Authentication Service

| Endpoint   | Method | Success Status | Auth Token Requirement |
| ---------- | ------ | -------------- | --------------------- |
| /          | GET    | 200            | NO                    |
| /register  | POST   | 201            | NO                    |
| /login     | POST   | 200            | NO                    |
| /logout    | POST   | 202            | YES                   |

***

## User Service
| Endpoint                    | Method | Success Status | Auth Token Requirement |
| --------------------------- | ------ | -------------- | --------------------- |
| /                           | GET    | 200            | NO                    |
| /delete                     | DELETE | 204            | YES                   |
| /follow?toFollowId=         | POST   | 200            | YES                   |
| /unfollow?toUnfollowId=     | POST   | 200            | YES                   |

***
## WebSocket Service

| Endpoint                    | Method | Success Status | Auth Token Requirement |
| --------------------------- | ------ | -------------- | --------------------- |
| /                           | GET    | 200            | NO                    |
| /ws/createRoom              | POST   | 200            | YES                   |
| /ws/joinRoom/{roomId}       | GET    | 200            | YES                   |
| /ws/getRooms                | GET    | 200            | YES                   |
| /ws/getClients/:roomId      | GET    | 200            | YES                   |

***
## Database Schema
![image](https://github.com/RipulHandoo/Go-Chat/assets/107461226/3e901df2-ed82-4b04-ab38-acc3f69dc6f2)
https://drawsql.app/teams/ripuls-team/diagrams/go-chat

## Datbase Description
| Table Name        | Description                                                                                     |
| ----------------- | ----------------------------------------------------------------------------------------------- |
| users             | Stores the details of users. The password is not stored directly; it is hashed for security purposes. |
| posts             | Stores the posts created by users.                                                        |
| user_followers    | Represents the relationship between users, where the follower_id follows the following_id.     |

***
# Web Socket Architecture
- # Hub Architecture
  
![hub_architecture](https://github.com/RipulHandoo/Go-Chat/assets/107461226/a03aedf4-801c-458c-9778-ae049ffeaf4c)

First, we have the hub running on a separate goroutine which is the central place that manages different channels and contains a map of rooms. The hub has a Register and an Unregister channel to register/unregister clients, and a Broadcast channel that receives a message and broadcasts it out to all the other clients in the same room.

![join_room](https://github.com/RipulHandoo/Go-Chat/assets/107461226/eef7c3f4-b188-425d-9e4d-9dea1cc70f87)

A room is initially empty. Only when a client hits the /ws/joinRoom endpoint, that will create a new client object in the room and it will be registered through the hub's Register channel.

![hub_initial](https://github.com/RipulHandoo/Go-Chat/assets/107461226/6c1702e9-479d-4872-a012-235348f41820)

Each client has a writeMessage and a readMessage method. readMessage reads the message through the client's websocket connection and send the message to the Broadcast channel in the hub, which will then broadcast the message out to every client in the same room. The writeMessage method in each of those clients will write the message to its websocket connection, which will be handled on the frontend side to display the messages accordingly.

***
## Tools and Technologies Used

- Go
- PostgreSQL
- Docker
