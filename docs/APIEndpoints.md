**WebScrapingChallenge API Description** 

The WebScrapingChallenge API offers a set of endpoints designed for web scraping and sentiment analysis: 
1. **Available Tags Endpoint** : 
- **Endpoint** : `/availableTags` 
- **HTTP Method** : `GET` 
- **Description** : Fetches a list of available HTML tags that can be used as filters during web scraping. 
- **Response** : Returns an array of HTML tags. 
2. **Scrape Endpoint** : 
- **Endpoint** : `/scrape` 
- **HTTP Method** : `POST` 
- **Description** : Scrapes the content of a given URL. Users can optionally filter the scraping by specific HTML tags. 
- **Parameters** : 
- `url` (required): The URL of the webpage to scrape. 
- `tags` (optional): List of HTML tags to filter by. If not provided, all available tags will be used. 
- **Response** : Provides the total word count of the scraped page and content segmented by specified HTML tags. Each segment contains all appearances of that HTML tag. 
3. **Sentiment Analysis Endpoint** : 
- **Endpoint** : `/analyzeSentiment` 
- **HTTP Method** : `POST` 
- **Description** : Analyzes the sentiment of the content at a given URL. 
- **Parameters** : 
- `url` (required): The URL of the webpage to analyze. 
- **Response** : Returns the sentiment of the webpage's content, which can be either `positive`, `negative`, or `neutral`.

By using this API, users can easily fetch content from web pages, segment it by HTML tags, and gain insight into the overall sentiment of the content.