<script>
  import Task from './Task.svelte';
  import SearchTask from './SearchTask.svelte';

  export let taskList;
</script>

<SearchTask/>

<table>
  <thead>
    <tr>
      <th>Name</th>
      <th>Description</th>
      <th>Due Date</th>
      <th>Creation Date</th>
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
  table {
    width: 100%;
  }

  table th {
    text-align: left;
  }
</style>