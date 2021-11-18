import fetch from "node-fetch";
var k = "Shortened url will be displayed here"
document.getElementById("a").innerText=k
/*
document.getElementById("submit").addEventListener("click",function(event){
      var long=document.getElementById("long").value.trim()
      const todo = {
        "longURL": long
        };
  fetch("https://myproxxyserver.herokuapp.com/https://url--shortener-server.herokuapp.com/shorten", {
    method: "POST",
    body: JSON.stringify(todo),
    headers: {"Content-type": "application/json; charset=UTF-8"}
  })
  .then(response => response.json()) 
  .then(json => console.log(json));
},)
*/