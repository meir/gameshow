
<script>
  import Start from "../components/start.svelte";
  import Game from "../components/game.svelte";
  $: user = null;
</script>

{#if user == null}
  <Start submit={(data) => user = data}/>
{:else}
  <Game me={user} />
{/if}

