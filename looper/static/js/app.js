let tasks = document.getElementById("tasks");
axios.get("http://localhost:8080/tasks").then((resp) => {
  tasks.innerHTML = "";
  for (let task of resp.data.tasks) {
    let div = document.createElement("div");

    let input = document.createElement("input");
    input.type = "checkbox";
    input.addEventListener("change", () => {
      tasks.removeChild(div);
    })
    div.appendChild(input);

    let text = document.createTextNode(" " + task);
    div.appendChild(text);

    tasks.appendChild(div);
  }
})
