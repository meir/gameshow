
<script>
  import Profile from "./profile.svelte";
  import Score from "./score.svelte";
  import { dndzone } from "svelte-dnd-action";

  export let team = [];
  export let hideComponents = false;

  const consider = (e) => {
    team = e.detail.items;
  };
  
  export const setTeam = (e) => {};
</script>

<style>
section {
  border: 5px solid #105353;
  border-radius: 15px;
  background-color: #f8e3e4;
  display: flex;
  justify-content: space-around;
  padding: 20px;
  margin-top: 5px;
}

div {
  margin: 20px;
}
</style>

<div>
  <Score />
  <section use:dndzone={{items: team, dragDisabled: !hideComponents}} on:consider={consider} on:finalize={setTeam}>
    {#each team as player(player.id)}
      <Profile img={player.picture} color={player.color} name={player.name} />
    {/each}
  </section>
</div>

