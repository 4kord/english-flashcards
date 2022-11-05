<button>
    <slot></slot>
</button>

<style>
    button {
        background-color: #682cd4;
        color: inherit;
        border: none;
        border-radius: 5px;
        padding: 10px;
        font: inherit;
    }
    button:hover {
        cursor: pointer;
        background-color: #9074f4;
    }
    button:active {
        border: 3px solid #4c3b81;
        padding: 7px;
        background-color: #271352;
        cursor: pointer;
    }
</style>