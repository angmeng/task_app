<script>
  import { GetTaskList } from './tasks.js';
  
  let searchText = "";

  function onSubmit() {
    GetTaskList(searchText, []);
  }

  function reset() {
    searchText = "";
    GetTaskList("", []);
  }
</script>

<div>
  <form>
    <input type="text" placeholder="Search by name..." name="search" bind:value={searchText}>
    <button on:click={onSubmit}>Search</button>
    <button on:click={reset}>Reset</button>
  </form>
</div>