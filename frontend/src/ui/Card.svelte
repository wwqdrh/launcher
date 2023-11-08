<script>
    import { Card } from "https://wwqdrh.github.io/assets/uikit/common.es.js";
    import { onMount } from "svelte";

    let el;
    onMount(() => {
        new Card({
            target: el,
        });
    });
</script>

<div bind:this={el} />
