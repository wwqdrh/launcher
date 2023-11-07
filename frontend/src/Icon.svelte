<script>
    export let Letter = "A";
    export let onClick = () => {};
</script>

<button on:click={onClick} class="icon">
    <span class="letter">{Letter.substring(0, 1) || "A"}</span>
</button>

<style>
    .icon {
        width: 128px;
        height: 128px;
        background-color: blue;
        border-radius: 50%;
        cursor: pointer;
        display: flex;
        align-items: center;
        justify-content: center;
    }

    .letter {
        color: white;
        text-transform: uppercase;
        font-size: 72px;
        font-family: Arial, sans-serif;
    }
</style>
