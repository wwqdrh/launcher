<script>
    export let title = "setting";

    export function toggle() {
        if(el$$) {
            el$$.toggle()
        }
    }

    import { Drawder } from "https://wwqdrh.github.io/assets/uikit/common.es.js";
    import { onMount } from "svelte";

    let el;
    let el$$;
    onMount(() => {
        el$$ = new Drawder({
            target: el,
            props: {
                title,
            },
        });
    });
</script>

<div bind:this={el} />
