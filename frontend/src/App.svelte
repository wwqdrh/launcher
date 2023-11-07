<script>
  import { AppList, StartApp } from "../wailsjs/go/main/App.js";
  import { onMount } from "svelte";

  import Icon from "./Icon.svelte";
  import {
    notify,
  } from "https://wwqdrh.github.io/assets/uikit/common.es.js";

  let appList = [];
  let iframeurl = "";

  let cont;

  function doAppList() {
    AppList().then((result) => (appList = [...result]));
  }

  function doStartApp(name) {
    StartApp(name).then((result) => {
      iframeurl = result;
    });
  }

  onMount(() => {
    doAppList();
    notify({ msg: "a test" });
    notify({
      type: "error",
      msg: "a error test asdjkalsdja dajskldjaklsd adsjkld",
    });
  });
</script>

<main>
  {#each appList as item}
    <div>
      <Icon onClick={() => doStartApp(item)} Letter={item} />
    </div>
  {/each}
  <div bind:this={cont} class="uikit-w-32 uikit-h-32 uikit-shadow-lg">
    ajdsklasdjlasdjakljdskjdlasdakljdka
  </div>
  <!-- <ContextMenu
    target={cont}
    menus={[
      {
        title: "item1",
        onClick: (v) => {
          let selection = window.getSelection();
          console.log(selection.toString());
        },
      },
      {
        title: "item2",
        onClick: (v) => {
          console.log(v);
        },
      },
    ]}
  /> -->
  {#if iframeurl}
    <div>
      <iframe title="app" width="400" height="400" src={iframeurl} />
    </div>
  {/if}
</main>

<style>
  @import url("https://wwqdrh.github.io/assets/uikit/common.css");
</style>
