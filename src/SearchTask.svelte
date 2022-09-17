<script>
  import { GetTaskList } from './tasks.js';
  
  let searchText = "";

  function search() {
    GetTaskList(searchText, []);
  }

  function reset() {
    searchText = "";
    GetTaskList("", []);
  }
</script>

<div>
  <input type="text" placeholder="Search by name..." name="search" bind:value={searchText}>
  <button on:click={search}>Search</button>
  <button on:click={reset}>Reset</button>
</div>