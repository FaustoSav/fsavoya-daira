<script>
	import '../routes/styles.css';

	export let btnInfo;
</script>

<button
	class="{btnInfo.customClass} btn"
	on:click={console.log('Click en el boton ', btnInfo.value)}
>
	{btnInfo.value}
</button>

<style>
</style>
