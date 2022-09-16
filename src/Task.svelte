<script>
  import { DeleteTask } from './tasks.js';

  export let task;

  function parseDate(date) {
    let d = new Date(date)
    return d.toLocaleDateString();
  }
</script>

<td>{task.name}</td>
<td>{task.description}</td>
<td>{parseDate(task.due_date)}</td>
<td>{parseDate(task.created_at)}</td>
<td>
  <a class="btn" href="#" on:click="{() => alert(task.id) }"> Edit</a> | 
  <a class="btn" href="#" on:click="{() => DeleteTask(task.id) }"> Delete</a>
</td>