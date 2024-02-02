<script>
	import '../routes/styles.css';

	let historyArray = ['2 - 2- 2- 2- 22 - 2- 2- 2- 22 - 2- 2- 2- 22 - 2- 2- 2- 22 - 2- 2- 2- 22 - 2- 2- 2- 2', '2 - 2', '2 - 2- 2- 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2', '2 - 2'];
</script>

<div class="history-container">
	<h1 class="text-lg uppercase text-gray-50 bg-[#2C2D34]w-full py-4 text-center font-mono tracking-widest">Historial</h1>
	<ul class="list-container">
		{#each historyArray as operation}
			<li class="list-item">{operation}</li>
		{/each}
	</ul>
	<h1 class="text-lg  text-gray-50 bg-[#2C2D34]w-full py-1 text-center uppercase text-[12px] font-mono tracking-widest">* Los datos no persisten en la base de datos</h1>
</div>

<style>
	.history-container {
		@apply w-[380px] h-[500px]  bg-opacity-30 rounded-3xl border-[1px] border-white border-opacity-10 overflow-hidden flex  flex-col justify-start items-center ;
	}
	.list-container {
		@apply  bg-black bg-opacity-30 w-full h-full text-gray-300 text-center overflow-y-scroll pt-1;
	}
	.list-item {
		@apply  p-0 border-b-[.5px] border-white border-opacity-5 px-5 py-1;
	}
</style>
