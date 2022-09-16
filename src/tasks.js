import { writable, derived } from 'svelte/store';

export const apiData = writable({});
export const query = writable({});

export const taskList = derived(apiData, ($taskResponse) => {
  if ($taskResponse.data){
    return $taskResponse.data.map(task => task);
  }
  return [];
});

export const meta = derived(apiData, ($taskResponse) => {
  if ($taskResponse.meta){
    return $taskResponse.meta;
  }
  return {};
});

export function GetTaskList(searchText) {
  console.log("search text: "+searchText)
  query["sort"] = ["+due_date", "+id"]
  if (searchText != "") {
    query["filter"] = {}
    query["filter"]["name"] = { "$like": "%"+searchText+"%" }
  } else {
    delete query["filter"]
  }

  let payload = JSON.stringify(query);
  // console.log(payload)
  let url = "http://localhost:3000/api/tasks?query="+window.btoa(payload);

  fetch(url)
      .then(response => response.json())
      .then(data => {
        // console.log(data);
        apiData.set(data);
      }).catch(error => {
        console.log(error);
        return [];
      });
}

export async function DeleteTask(id) {
  var result = confirm("confirm to delete?");

  if (result) {
    const res = await fetch('http://localhost:3000/api/tasks/' + id, {
			method: 'DELETE'
		})
		
		const json = await res.json()
		let result = JSON.stringify(json)
    if (res.status == 200) {
      alert("Task deleted successfully");
      GetTaskList("");
    } else {
      alert(result.message);
    }
  }
}

export async function UpdateTask(id, data) {
  var headers = new Headers();
    // uncomment below to attach the user access token to the header for authentication. 
    // headers.append("Authorization", "Basic YXBwbGU6MTIzNDU2Nzg=");
    headers.append("Content-Type", "application/json");
    const res = await fetch('http://localhost:3000/api/tasks/'+id, {
			method: 'PUT',
      headers: headers,
			body: JSON.stringify(data)
		})
		
		const json = await res.json()
		let result = JSON.stringify(json)
    if (res.status == 200) {
      alert("Task updated successfully");
      GetTaskList("");
    } else {
      alert(result.message);
    }
}