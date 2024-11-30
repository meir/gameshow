
<script>
  import Profile from "./profile.svelte";

  $: image = "https://avatars.githubusercontent.com/u/1"
  $: color = "#f00"
  $: name = "Mark"

  const uploadImage = () => {
    const file = document.getElementById("picture").files[0];
    const reader = new FileReader();
    reader.onload = () => {
      image = reader.result;
    };
    reader.readAsDataURL(file);
  };

  const changeColor = e => {
    color = e.target.value;
    console.log(color);
  };

  const changeName = e => {
    name = e.target.value;
    console.log(name);
  };

  export let submit = (data) => {};
  const submitData = () => {
    submit({
      id: Math.random(), // dont care
      name: name,
      color: color,
      picture: image,
    });
  };
</script>

<style>
main {
  width: 35%;
  height: 40%;
  margin: 35px;
  background-color: #f8e3e4;
  border: 5px solid #105353;
  border-radius: 15px;
  margin-top: 50px;
}

#preview {
  width: 100%;
  display: flex;
  justify-content: center;
  margin-top: 20px;
  margin-bottom: 20px;
}

label {
  display: block;
  text-transform: lowercase;
  opacity: 0.3;
  margin-left: 15px;
}

#name-block {
  width: 80%;
  margin-left: 10%;
  margin-bottom: 20px;
}

input[type="text"] {
  width: 100%;
  border: 4px solid #105353;
  border-radius: 15px;
  height: 40px;
  background-color: white;
  font-size: 20px;
  font-weight: bold;
}

input[type="color"] {
	-webkit-appearance: none;
	border: none;
	width: 200px;
	height: 35px;
  cursor: pointer;
}
input[type="color"]::-webkit-color-swatch-wrapper {
	padding: 0;

}
input[type="color"]::-webkit-color-swatch {
  border-radius: 15px;
  border: 4px solid #105353;
}

input[type="file"] {
  display: none;
}

#picture-label {
  cursor: pointer;
  border: 4px solid #105353;
  height: 30px;
  width: 200px;
  border-radius: 15px;
  background-color: #127373;
  font-size: 15px;
  justify-content: center;
  line-height: 28px;
  text-align: center;
  font-weight: bold;
  align-items: center;
  justify-content: center;
  opacity: 1;
  color: white;
  margin-left: 0;
}

button {
  width: 80%;
  margin-left: 10%;
  margin-right: 10%;
  border: 4px solid #105353;
  border-radius: 15px;
  height: 40px;
  background-color: #127373;
  color: white;
  font-size: 20px;
  font-weight: bold;
  text-transform: uppercase;
  margin-top: 20px;
  cursor: pointer;
}

.flex {
  display: flex;
  width: 80%;
  margin-left: 10%;
  justify-content: space-around;
}
</style>

<main>
  <div id="preview">
    <Profile img="{image}" color="{color}" name="{name}" />
  </div>

  <div id="name-block">
    <label for="name">Name</label>
    <input type="text" id="name" name="name" on:change={changeName}/>
  </div>

  <div class="flex">
    <div>
      <label for="picture">Picture</label>
      <label for="picture" id="picture-label">Choose a picture</label>
      <input type="file" accept="image/*" id="picture" name="picture" on:change={uploadImage} />
    </div>

    <div>
      <label for="color">Color</label>
      <input type="color" id="color" name="color" on:change={changeColor}/>
    </div>
  </div>

  <button on:click={submitData}>Join</button>
</main>
