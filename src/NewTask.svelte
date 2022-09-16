<script>
  let newTask = {"name": "", "description": "", "due_date": null};

  function addToTaskList() {
    alert("added successfully");
  }
</script>

<div>
  <label>
    Task name<br>
    <input bind:value={newTask.name} type="text" size="50" placeholder="what would you like to do...">
  </label>
  
  <label>
    Describe your task<br>
    <textarea bind:value={newTask.description} rows="4"  cols="50"></textarea>
  </label>
  
  <label>
    Due Date<br>
    <input type="date" required pattern="\d{4}-\d{2}-\d{2}" />
  <span class="validity"></span>
  </label>
  <button on:click={addToTaskList}>Add</button>
</div>

<style>
  div { padding: 20px 0px;}
</style>
