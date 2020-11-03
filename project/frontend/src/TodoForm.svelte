<script>
	export let btnName;
	export let txtPlaceholder;
</script>

<style>
	input[type="text"] {
		width: 50%;
		font-weight: 50;
	}
	input[type="submit"] {
		color: black;
		background-color: gold;
		text-transform: uppercase;
		font-weight: 50;
	}
	main {
		text-align: center;
	}
	[id="add-todo"] {
		align-content: center;
	}
	[id="todo-p"] {
		text-align: center;
	}
</style>

<main>
	<div id="add-todo">
		<form action="/todos" method="post">
			<input
				type="text"
				placeholder={txtPlaceholder}
				id={btnName}
				name={btnName}
				required
				minlength="2"
				maxlength="140" />
			<input type="submit" value={btnName} />
		</form>
	</div>
</main>
