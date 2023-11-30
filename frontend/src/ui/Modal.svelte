<script>
    export let handle;
    export let content;

    import { Modal } from "https://wwqdrh.github.io/assets/uikit/common.es.js";
    import { onMount } from "svelte";

    let el;
    onMount(() => {
        new Modal({
            target: el,
            props: {
                handle,
                content,
            }
        });
    });
</script>

<div bind:this={el} />
