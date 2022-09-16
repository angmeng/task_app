import { writable, derived } from 'svelte/store';

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

export function GetTaskList() {
  fetch("http://localhost:3000/api/tasks")
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
      GetTaskList();
    } else {
      alert(result.message);
    }
  }
}