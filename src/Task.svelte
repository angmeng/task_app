<script>
  import { DeleteTask, UpdateTask } from './tasks.js';

  export let task;
  let editable = false;

  function parseDate(date) {
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
</script>

{#if editable == false}
  <td>{task.name}</td>
  <td>{task.description}</td>
  <td>{parseDate(task.due_date)}</td>
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
<td>{parseDate(task.created_at)}</td>
<td>
  {#if task.status_id == 0}
    Not Urgent
  {:else if task.status_id == 1}
    Due Soon
  {:else if task.status_id == 2}
    Overdue
  {:else}
    Unknown
  {/if}
</td>
<td>
  {#if editable == false}
    <a href="#" on:click="{() => editable = !editable }"> Edit</a> | 
    <a href="#" on:click="{() => DeleteTask(task.id) }"> Delete</a>
  {:else}
    <button on:click={ updateTask }> Update</button>
  {/if}
</td>