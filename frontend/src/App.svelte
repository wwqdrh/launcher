<script>
  import { AppList, StartApp } from "../wailsjs/go/main/App.js";
  import { onMount } from "svelte";

  import Icon from "./ui/Icon.svelte";
  import {
    notify,
    Modal as IModal,
  } from "https://wwqdrh.github.io/assets/uikit/common.es.js";
  import Card from "./ui/Card.svelte";
  import Modal from "./ui/Modal.svelte";

  let appList = [];
  let iframeurl = "";

  let cont;
  let openmodal;

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
    new IModal({
      target: openmodal,
      props: {
        handle: document.getElementById("customhandle"),
        content: (() => {
          let dom = document.createElement("div");
          dom.innerHTML = `<p>2</p>`;
          return dom;
        })(),
      },
    });
  });
</script>

<main>
  {#each appList as item}
    <div>
      <Icon
        onClick={() => document.getElementById("customhandle").click()}
        Letter={item}
      />
      <Card />
    </div>
  {/each}
  <div id="customhandle" />
  <div bind:this={openmodal}>1</div>
  {#if iframeurl}
    <div>
      <iframe title="app" width="400" height="400" src={iframeurl} />
    </div>
  {/if}
</main>

<style>
  @import url("https://wwqdrh.github.io/assets/uikit/common.css");
</style>
