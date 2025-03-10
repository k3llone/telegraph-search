
const searchButton = document.querySelector('#search-button');
const searchLine = document.querySelector('#search-line')
let searchCount = 0;
const searchResultArr = [];

searchButton.addEventListener('click', function() {
    const searchResults = document.getElementById("search-results")

    console.log(searchLine.value)

    runSearch(searchLine.value)
});

function runSearch(text) {
    const socket = new WebSocket("ws://" + location.host + "/ws");

    socket.onopen = function(event) {
        socket.send(text)
    };
    
    socket.onmessage = function(event) {
        console.log("Message from server:", event.data);
        var obj = JSON.parse(event.data);
        searchCount += 1; 
        if (obj.status != 404){
            addResult(searchCount, obj.title, obj.videos.length, obj.links.length, obj.images.length, obj.size.toFixed(2), obj.link)
            searchResultArr.push(obj)
        }
    };
    
    socket.onerror = function(error) {
        console.error("WebSocket error:", error);
    };
    
    socket.onclose = function(event) {
        console.log("WebSocket connection closed.");
        console.log(searchCount)
    };
}

function addResult(id, title, videos, links, images, size, link, ) {
    const searcResults = document.getElementById("search-results")

    const resultTitle = document.createElement("h1")
    resultTitle.className = "result-title"
    resultTitle.textContent = title

    const resultTagVideos = document.createElement("div")
    resultTagVideos.className = "tag-video"
    resultTagVideos.append(document.createElement("span").textContent = '🎬')
    resultTagVideos.append(document.createElement("span").textContent = ` (${videos})`)

    const resultTagLinks = document.createElement("div")
    resultTagLinks.className = "tag-links"
    resultTagLinks.append(document.createElement("span").textContent = '🔗')
    resultTagLinks.append(document.createElement("span").textContent = ` (${links})`)

    const resultTagImages = document.createElement("div")
    resultTagImages.className = "tag-images"
    resultTagImages.append(document.createElement("span").textContent = '🖼')
    resultTagImages.append(document.createElement("span").textContent = ` (${images})`)

    const resultTagSize = document.createElement("div")
    resultTagSize.className = "tag-size"
    resultTagSize.append(document.createElement("span").textContent = '📦')
    resultTagSize.append(document.createElement("span").textContent = ` ${size}Kb`)

    const resultTags = document.createElement("div")
    resultTags.className = "result-tags"

    if (videos > 0){
        resultTags.append(resultTagVideos)
    }
    if (links > 0){
        resultTags.append(resultTagLinks)
    }
    if (images > 0) {
        resultTags.append(resultTagImages)
    }
    if (size > 0) {
        resultTags.append(resultTagSize)
    }

    const resultOpenButton = document.createElement("input")
    resultOpenButton.className = "result-button-open"
    resultOpenButton.type = "button"
    resultOpenButton.value = "Open"
    resultOpenButton.id = id

    resultOpenButton.addEventListener('click', function() {
        window.open(link, '_blank').focus();
    });

    const resultLinkButton = document.createElement("input")
    resultLinkButton.className = "result-button-link"
    resultLinkButton.type = "button"
    resultLinkButton.value = "Link"
    resultLinkButton.id = id

    resultLinkButton.addEventListener('click', function() {
        navigator.clipboard.writeText(link).then(function() {
            console.log('Finaly');
          }, function(err) {
            console.error('Error: ', err);
          });
    });

    const result = document.createElement("div")
    result.className = "search-result"
    result.id = id

    result.append(resultTitle)
    result.append(resultTags)
    result.append(resultOpenButton)
    result.append(resultLinkButton)


    searcResults.append(result)
}