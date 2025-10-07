const axios = require('axios');

const baseUrl = 'https://search.censys.io/api/v2/hosts/search';
const headers = {
  'Accept': 'application/json',
  'Authorization': 'Basic [BASE64 KEY]'
};

function sendRequest(url) {
    axios.get(url, { headers })
        .then(response => {
        const { data } = response;
        const { result } = data;
        const { hits } = result;

        for (const hit of hits) {
            console.log(`${hit.ip}`);
        }

        if (result.links && result.links.next) {
            const nextUrl = `${url}&cursor=${result.links.next}`;
            sendRequest(nextUrl);
        }
    })
    .catch(error => {
        console.error('Error:', error.message);
    });
}

const initialUrl = `${baseUrl}?q=%28FlareSolverr%29%20and%20services.service_name%3D%60HTTP%60&per_page=100&virtual_hosts=EXCLUDE`;
sendRequest(initialUrl);
