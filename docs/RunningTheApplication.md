**Running the WebScrapingChallenge Application**  
1. **Clone the Repository** :
- Open a terminal or command prompt.
- Navigate to the directory where you want to clone the repository. 
- Run the following command:

```bash

git clone https://github.com/AlexandruRusmir/WebScrapingChallenge.git
``` 
2. **Setting up the API** :

a. **Navigate to the API Directory** :

```bash

cd WebScrapingChallenge/app/api
```



b. **Install Dependencies** : 
- Ensure you have Go installed on your system. If not, [download and install it](https://golang.org/dl/) .
- Since Go dependencies are typically included in the codebase, you may not need to install additional packages. If needed, dependencies will be fetched automatically when you run the app.

c. **Run the API** :

```go

go run cmd/main.go
```



This will start the API on `localhost` at port `4200`. 

3. **Setting up the Vue.js UI** :

a. **Navigate to the UI Directory** :

```bash

cd ../ui
```



b. **Install Dependencies** : 
- Ensure you have Node.js and npm installed. If not, [download and install them](https://nodejs.org/) . 
- Install the required npm packages:

```

npm install
```

c. **Run the Vue App** :

```arduino

npm run serve
```



Once compiled successfully, you should be able to access the Vue app on your browser, typically at `localhost:8080`.

4. **Access the Application** :
- With both the API and Vue app running, open your web browser. 
- Navigate to the Vue app URL, which is usually `http://localhost:8080`. 
- You can now interact with the application, and it will communicate with the API running on port `4200`.