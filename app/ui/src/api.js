const BASE_URL = 'http://localhost:4200';

/**
 * @returns {Promise<string[]>} - The available tags
 */
const fetchAvailableTags = async () => {
    const endpoint = `${BASE_URL}/availableTags`;
    const response = await fetch(endpoint);
    return await response.json();
}

/**
 * Send a request to scrape content
 * @param {string} url - The URL to be scraped
 * @param {string[]} [tags] - An optional array of HTML tags to scrape
 * @returns {Promise<object>} - The scraped data
 */
const scrapeContent = async (url, tags) => {
    const endpoint = `${BASE_URL}/scrape`;
    const body = {
        url: url,
    };
    if (tags && tags.length > 0) {
        body.tags = tags;
    }

    const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    });

    return await response.json();
}

/**
 * Send a request to analyze sentiment
 * @param {string} url - The URL to be analyzed
 * @returns {Promise<string>} - The sentiment analysis result
 */
const analyzeSentiment = async (url) => {
    const endpoint = `${BASE_URL}/analyzeSentiment`;
    const body = {
        url: url
    };

    const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(body)
    });

    return await response.json();
}

export {
    scrapeContent,
    analyzeSentiment,
    fetchAvailableTags
};