
<script>
  import Profile from "./profile.svelte";
  import { dndzone } from "svelte-dnd-action";

  export let teams = [[]];
  export let team = 0;
  export let isHost = false;

  const considerDnd = (e) => {
    teams[team] = e.detail.items;
  };

  const finalizeDnd = (e) => {
    teams[team] = e.detail.items;
  };
</script>

<style>
section {
  border: 5px solid #105353;
  border-radius: 15px;
  background-color: #f8e3e4;
  display: flex;
  justify-content: space-around;
  padding: 20px;
  margin: 20px;
}
</style>

<section use:dndzone={{items: teams[team], dragDisabled: !isHost}} on:consider={considerDnd} on:finalize={finalizeDnd}>
  {#each teams[team] as player(player.name)}
    <Profile img={player.picture} color={player.color} name={player.name} />
  {/each}
</section>
