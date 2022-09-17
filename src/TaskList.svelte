<script>
  import Task from './Task.svelte';
  import SearchTask from './SearchTask.svelte';
  import {GetTaskList} from './tasks.js';

  export let taskList;
  let dueDateSortingImg = "/ascending.png";
  let idSortingImg = "/ascending.png";
  let sortData = ["+due_date", "+created_at"];

  function updateSortingIcon(subject) {
    if (subject == "due_date") {
      if (sortData[0] == "+due_date") {
        sortData = ["-due_date", sortData[1]];
        dueDateSortingImg = "/descending.png"
      } else if (sortData[0] == "-due_date") {
        sortData = ["+due_date", sortData[1]];
        dueDateSortingImg = "/ascending.png"
      }  else if (sortData[1] == "+due_date") {
        sortData = ["-due_date", sortData[0]];
        dueDateSortingImg = "/descending.png"
      } else if (sortData[1] == "-due_date") {
        sortData = ["+due_date", sortData[0]];
        dueDateSortingImg = "/ascending.png"
      }
    } else if (subject == "created_at") {
      if (sortData[0] == "+created_at") {
        sortData = ["-created_at", sortData[1]];
        idSortingImg = "/descending.png"
      } else if (sortData[0] == "-created_at") {
        sortData = ["+created_at", sortData[1]];
        idSortingImg = "/ascending.png"
      } else if (sortData[1] == "+created_at") {
        sortData = ["-created_at", sortData[0]];
        idSortingImg = "/descending.png"
      } else if (sortData[1] == "-created_at") {
        sortData = ["+created_at", sortData[0]];
        idSortingImg = "/ascending.png"
      }
    }

    GetTaskList("", sortData);
  }
</script>

<SearchTask/>

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Description</th>
      <th>Due Date <img src={dueDateSortingImg} on:click="{() => updateSortingIcon("due_date")}"></th>
      <th>Creation Date <img src={idSortingImg} on:click="{() => updateSortingIcon("created_at")}"></th>
      <th>Status</th>
      <th>Actions</th>
    </tr>
  </thead>
  <tbody>
    {#each taskList as task}
      <tr>
        <Task task={task}/>
      </tr>
    {/each}
  </tbody>
</table>

<style>
  table { width: 100%; }
  table th { text-align: left; }
  img { width: 25px; padding-left: 5px; }
</style>