<script>
	import '../routes/styles.css';
	import { getOperationsHistoy } from '../utils/Get';
	import ReloadIcon from './ReloadIcon.svelte';
	import { onMount } from 'svelte';
	let HistoryItems = [];

	onMount(() => {
		// Tu lógica o función que quieres ejecutar al montar el componente
		getHistory();
	});

	const getHistory = () => {
		getOperationsHistoy()
			.then((result) => {
				HistoryItems = result;
				console.log(HistoryItems);
			})
			.catch((error) => console.error(error));
	};
</script>

<div class="history-container">
	<h2 class="title w-full">Historial <ReloadIcon onclickProp={getHistory} /></h2>

	<ul class="list-container">
		{#if HistoryItems.length > 0}
			{#each HistoryItems as operation (operation.id)}
				<li class="list-item hover:bg-[#2C2D34]">{operation.operation}</li>
			{/each}
		{:else}
			<li class="empty-history">No hay datos</li>
		{/if}
	</ul>
	<div class="  w-full h-[36px] font-mono tracking-widest"></div>
</div>

<style>
	.history-container {
		@apply w-[380px] h-[500px]  bg-opacity-30 rounded-3xl border-[1px] border-white border-opacity-10 overflow-hidden flex  flex-col justify-start items-center;
	}
	.list-container {
		@apply bg-black bg-opacity-30 w-full h-full text-gray-300 text-center overflow-y-scroll;
	}
	.list-item {
		@apply p-0 border-b-[.5px] border-white border-opacity-5 px-5 pt-1;
	}
	.empty-history {
		@apply w-full h-full  bg-black bg-opacity-30 text-center text-xl uppercase font-mono pt-20;
	}
	.title {
		@apply text-lg uppercase text-gray-50 w-full py-1 text-center font-mono tracking-widest relative;
	}
</style>
