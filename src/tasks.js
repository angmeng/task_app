import { debug } from 'svelte/internal';
import { writable, derived } from 'svelte/store';

let query = {};
let sortData = []
export const apiData = writable({});

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

export function GetTaskList(searchText, sorting) {
  if (searchText != "") {
    query["filter"] = {}
    query["filter"]["name"] = { "$like": "%"+searchText+"%" }
  } else {
    delete query["filter"]
  }

  if (sortData.length == 0) {
    sortData = ["+due_date", "+created_at"];
  }

  if (sorting.length > 0) {
    sortData = sorting;
  }

  query["sort"] = sortData

  let payload = JSON.stringify(query);
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
      GetTaskList("", []);
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
      GetTaskList("", []);
    } else {
      alert(result.message);
    }
}