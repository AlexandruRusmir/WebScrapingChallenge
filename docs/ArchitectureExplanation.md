## Application Architecture Overview

The WebScrapingChallenge application is structured into two main components: a backend API and a frontend user interface.
### Backend (API):

The API is built using **Go**  , a choice inspired by my desire to venture into an unfamiliar language. The core structure of the API is as follows: 
1. **cmd Directory**  : 
- Contains `main.go`, which acts as the entry point of the API. Here, all the routes for our endpoints are defined and linked with their respective functionalities. 
2. **Middlewares**  : 
- Contains `cors.go` which handles Cross-Origin Resource Sharing (CORS), ensuring that the API is accessible to the frontend application. 
3. **Handlers**  :
- Responsible for managing HTTP requests to the different endpoints. It bridges the gap between incoming requests and the application's business logic. 
4. **Services**  : 
- Comprises the heart of the API. The services for scraping and sentiment analysis encapsulate the business logic required for the endpoints' operations, making use of modular functions. 
    - **Scraping Service** :
        - This service manages the extraction of page content and computes the word count. 
        - Special handling has been integrated for `a` and `img` elements: 
            - For `a` elements, both the `href` and the text content are displayed. 
            - For `img` elements, the `src` (source URL) and the alt text are presented.
        - For other HTML elements, their content is directly displayed, with elements being categorically grouped based on their type. 
    - **Sentiment Analysis Service** :
        - This service tokenizes the entire page content.
        - A score is computed based on the words encountered, referencing a list of predefined positive and negative words.
        - Using the derived score, the overall sentiment of the page (positive, negative, or neutral) is determined and returned.
5. **Utils**  :
- Equips the API with a set of utilities that aid in fetching data from URLs and manipulating text data. 
6. **Packages**  : 
- The application utilizes the `goquery` package which offers powerful scraping capabilities, allowing the extraction and manipulation of HTML content with ease

### Frontend:

For the frontend, **Vue.js**  was the chosen framework. My choice was influenced again by the curiosity to delve into an unfamiliar territory. Here's an insight into its structure: 
1. **App.vue** :
- The central component of our Vue application, from which several sub-components branch out to shape the entirety of our page. 
2. **Subcomponents** :
- These modular Vue components handle various functionalities of the frontend, from input forms to result displays. 
3. **API Interaction** :
- A designated file facilitates interactions between the frontend and the backend, making API calls and handling responses. 
4. **Unique Feature** :
- One distinct feature of the frontend is the ability for users to download the results of a page analysis in a JSON format locally, providing a tangible and portable outcome from their analysis.

### Reflection :

Embarking on this project was exciting, because it provided an opportunity to explore new technologies in order to broaden horizons. Particularly, the power and efficiency of Go stood out, marking it as a potential favorite for future projects. It's a journey that began with curiosity and culminated in a rich learning experience. I eagerly anticipate deepening my understanding of these and other new technologies in the future.