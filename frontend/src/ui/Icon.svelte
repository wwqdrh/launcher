<script>
    export let onClick = () => {};
    export let Letter = "";

    import { CircleIcon } from "https://wwqdrh.github.io/assets/uikit/common.es.js";
    import { onMount } from "svelte";

    let el;
    onMount(() => {
        new CircleIcon({
            target: el,
            props: {
                onClick,
                Letter
            },
        });
    });
</script>

<div bind:this={el} />
