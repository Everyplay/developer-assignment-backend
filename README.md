# developer-assignment-backend

Return your answer as a zip file containing all relevant files. 

## Software design

Use at most 2 hours total when answering these questions, so no need to go into minor details.

We want to store game specific data for users. Design a backend service that stores this data and supports the following requirements & use cases:

- We have 10 000 mobile games which have on average 10 000 daily active users, the service must be designed to support this load
- Store users data in a game
  - This data can be different for each game and may contain any number of properties e.g. level, score, etc.
  - Data may be updated multiple times per game session
- Game developer can query user data with following constraints:
  - Property is equal to given value
  - Property's value is less/greater than given value
  - Multiple constraints may be included in the query
  - For example: "give me user ids in my game who are in level 2 who have score greater than 9000"


### 1. Service architecture
Describe the overall architecture of the service, what components are needed and what is their responsibility. What things to consider when designing a service with this kind of scale? Focus on the software components, no need to think about hardware etc.

### 2. Data storage
Design the data storage for this service. Describe what database(s) to use and model the data. Use diagrams, provide the database schemas or use other relevant methods to describe the data model. Explain the pros/cons of your solution.

### 3. Design the API for the service
Describe the backend's API(s) to support the given use cases.

## Programming task

Implement a message hub.