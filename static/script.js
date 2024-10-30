let debounceTimer;
async function fetchSuggestions() {
    clearTimeout(debounceTimer);
    debounceTimer = setTimeout(async () => {
        const query = document.getElementById('search-input').value;
        if (query.length < 1) {
            document.getElementById('suggestions').innerHTML = '';
            return;
        }
        const response = await fetch(`/search?query=${query}`);
        const suggestions = await response.json();
        const suggestionsList = document.getElementById('suggestions');
        suggestionsList.innerHTML = '';

        suggestions.forEach(suggestion => {
            const li = document.createElement('li');
            const a = document.createElement('a');

            if (suggestion[0] === "-1") {
                li.textContent = suggestion[1];
            } else {
                if (suggestion[1].includes(" - location")) {
                    a.href = `/locations?id=${suggestion[0]}`; 
                } else {
                    a.href = `/details?id=${suggestion[0]}`; 
                }
                a.textContent = suggestion[1]; 
            }

            li.appendChild(a);
            suggestionsList.appendChild(li);
        });
        
    }, 300); // Delay time (ms)
}