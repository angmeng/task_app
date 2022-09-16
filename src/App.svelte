<script>
	import { onMount } from "svelte";
  import { apiData, taskList, meta } from './tasks.js';
	import TaskList from './TaskList.svelte';
	import Pagination from './Pagination.svelte';
	import NewTask from './NewTask.svelte';

  onMount(async () => {
		fetch("http://localhost:3000/api/tasks")
		.then(response => response.json())
		.then(data => {
			console.log(data);
			apiData.set(data);
		}).catch(error => {
			console.log(error);
			return [];
		});
	});
</script>

<main>
	<h1>Task App</h1>
	<NewTask/>
	<TaskList taskList={taskList}/>
	<Pagination meta={$meta}/>
</main>

<style>
	main {
		/* text-align: center; */
		padding: 1em;
		max-width: 240px;
		margin: 0 auto;
	}

	h1 {
		color: #ff3e00;
		text-transform: uppercase;
		font-size: 4em;
		font-weight: 100;
	}

	@media (min-width: 640px) {
		main {
			max-width: none;
		}
	}
</style>