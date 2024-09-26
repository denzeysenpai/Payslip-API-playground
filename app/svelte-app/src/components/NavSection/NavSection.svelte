<script lang="ts">
    import "./style.css"
    import Dash from "../../lib/dashboard.png"
    import CLD from "../../lib/calendar.png"
    import PS from "../../lib/payslip.png"
    import MNG from "../../lib/project-manager.png"
    let anchors = [
        {
            content: "Dashboard",
            path: "/dashboard",
            icon: Dash
        },
        {
            content: "Calendar",
            path: "/calendar",
            icon: CLD
        },
        {
            content: "Payslip",
            path: "/payslip",
            icon: PS
        },
        {
            content: "Manage",
            path: "/manage",
            icon: MNG
        }
    ]
</script>


<nav id="nav-wrapper">
    <div id="anchors-wrapper">
        {#each anchors as a}
            <a href="#{a.path}" class="anchor-nav" draggable="false">
                <img src="{a.icon}" alt="" draggable="false">
                {a.content}
            </a>
        {/each}
    </div>
</nav>