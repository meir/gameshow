
<script>
  export let img;
  export let color;
  export let name;
</script>

<style>
.profile {
  width: 90px;
}

.profile img {
  width: 80px;
  height: 80px;
  border-radius: 100%;
  border: 5px solid var(--color);
}

.profile div {
  width: 100%;
  text-align: center;
  color: var(--color);
}

</style>

<div class="profile" style="--color: {color}">
  <img src={img} alt={name} />
  <div>{name}</div>
</div>
