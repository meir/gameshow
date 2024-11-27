
<script>
  import Team from "./team.svelte";

  export let me = {};
  export let isHost = true;
  export let teams = [[me]];

  const addTeam = () => {
    teams = [...teams, []];
  };

  const removeTeam = () => {
    const len = teams.length;
    if (len < 2) {
      return;
    }

    const last = teams[len - 1];
    teams[len - 2] = [...teams[len - 2], ...last];
    teams = teams.slice(0, len - 1);
  };
</script>

<style>
.host {
  display: var(--is-host);
}

#team-container {
  display: flex;
  width: 95%;
  border: 5px solid #105353;
  border-radius: 15px;
  margin-top: 20px;
  margin-left: 2.5%;
  margin-right: 2.5%;
}

#teams {
  display: flex;
  width: 90%;
  justify-content: space-around;
}

#team-controls button {
  width: 50px;
  height: 50px;
  display: block;
  font-size: 30px;
  border: 5px solid #105353;
  border-radius: 100%;
  background-color: #f8e3e4;
  margin-top: auto;
  margin-bottom: auto;
  margin: 10px;
}
</style>

<div id="team-container" style="--is-host: {isHost ? "unset" : "none"}">
  <div id="teams">
    {#each teams as _, index}
      <Team isHost={isHost} teams={teams} team={index}/>
    {/each}
  </div>
  <div id="team-controls" class="host">
    <button on:click={addTeam}>+</button>
    <button on:click={removeTeam}>-</button>
  </div>
</div>

