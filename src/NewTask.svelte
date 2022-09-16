<script>
  import { GetTaskList } from './tasks.js';

  let result = null;

  async function onSubmit(e) {
    const formData = new FormData(e.target);
    const data = {};
    for (let field of formData) {
      const [key, value] = field;
      if (key == "due_date") {
        data[key] = value + "T00:00:00Z"
      } else {
        data[key] = value;
      }
    }
    // console.log(JSON.stringify(data))

    var headers = new Headers();
    // uncomment below to attach the user access token to the header for authentication. 
    // headers.append("Authorization", "Basic YXBwbGU6MTIzNDU2Nzg=");
    headers.append("Content-Type", "application/json");
    const res = await fetch('http://localhost:3000/api/tasks', {
			method: 'POST',
      headers: headers,
			body: JSON.stringify(data)
		})
		
		const json = await res.json()
		result = JSON.stringify(json)
    if (res.status == 201) {
      alert("Task added successfully");
      e.target.reset();
      GetTaskList("");
    } else {
      alert(result.message);
    }
  }
</script>

<div>
  <form on:submit|preventDefault={onSubmit}>
    <label>
      Task name<br>
      <input name="name" type="text" size="50" required placeholder="what would you like to do...">
    </label>
    
    <label>
      Describe your task<br>
      <textarea name="description" rows="4" cols="50" required></textarea>
    </label>
    
    <label>
      Due Date<br>
      <input type="date" name="due_date" required pattern="\d{4}-\d{2}-\d{2}" />
      <span class="validity"></span>
    </label>
    
    <button type="submit">Add to task list</button>
  </form>
</div>

<style>
  div { padding: 20px 0px;}
</style>
