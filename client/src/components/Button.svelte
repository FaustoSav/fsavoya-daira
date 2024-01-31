<script>
	import '../routes/styles.css'

  export let value, customClass;
</script>

<button class="{customClass} btn" on:click={console.log("Click en el boton ",value)}>
  {value}
</button>

<style>
	
</style>
