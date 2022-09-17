<script>
  import { DeleteTask, UpdateTask } from './tasks.js';

  export let task;
  let editable = false;
  let hourMs = 1000 * 60 * 60;
  let sevenDaysMs = hourMs * 24 * 7

  function parseDateToString(date) {
    let d = new Date(date)
    return d.toLocaleDateString();
  }

  function updateTask() {
    const data = {};
    data["name"] = task.name;
    data["description"] = task.description;
    data["due_date"] = task.due_date;
    UpdateTask(task.id, data);
    editable = !editable;
  }

  function getStatus() {
    let dueDate = new Date(task.due_date);
    let currentDate = new Date();
    let diffMs = dueDate.getTime() - currentDate.getTime();
    if (diffMs > 0) {
      if (diffMs > sevenDaysMs) {
        return "Not Urgent"
      } else {
        return "Due Soon"
      }
    } else {
      return "Overdue"
    }
  }
</script>

{#if editable == false}
  <td>{task.name}</td>
  <td>{task.description}</td>
  <td>{parseDateToString(task.due_date)}</td>
{:else}
  <td>
    <input name="name" type="text" size="20" required placeholder="what would you like to do..." bind:value={task.name}>
  </td>
  <td>
    <input name="description" type="text" size="50" required bind:value={task.description}>
  </td>
  <td>
    <input name="due_date" type="text" size="20" required bind:value={task.due_date}>
  </td>
{/if}
<td>{parseDateToString(task.created_at)}</td>
<td>
  {getStatus()}
</td>
<td>
  {#if editable == false}
    <a href="#" on:click="{() => editable = !editable }"> Edit</a> | 
    <a href="#" on:click={ DeleteTask(task.id) }> Delete</a>
  {:else}
    <button on:click={ updateTask }> Update</button> | 
    <a href="#" on:click="{() => editable = !editable }"> Cancel</a>
  {/if}
</td>