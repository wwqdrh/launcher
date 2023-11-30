<script>
    export let icon = "";
    export let home = "/";
    export let menus = [];
    export let onClick = (v) => {};

    import { Header } from "https://wwqdrh.github.io/assets/uikit/common.es.js";
    import { onMount } from "svelte";

    let el;
    onMount(() => {
        new Header({
            target: el,
            props: {
                icon,
                home,
                menus,
                onClick,
            },
        });
    });
</script>

<div bind:this={el} />
