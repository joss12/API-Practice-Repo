import mermaid from "https://cdn.jsdelivr.net/npm/mermaid@10/dist/mermaid.esm.min.mjs";
mermaid.initialize({ startOnLoad: false });

const typeSelect = document.getElementById("type");
const nodeFields = document.getElementById("node-fields");
const edgeFields = document.getElementById("edge-fields");

typeSelect.onchange = () => {
  if (typeSelect.value === "node") {
    nodeFields.style.display = "block";
    edgeFields.style.display = "none";
  } else {
    nodeFields.style.display = "none";
    edgeFields.style.display = "block";
  }
};

const form = document.getElementById("form");
const mermaidDiv = document.getElementById("mermaid");
const messageDiv = document.getElementById("message");

async function fetchGraphAndRender() {
  const res = await fetch("/map");
  const data = await res.json();

  const idMap = new Map(data.nodes.map((n) => [n.id, n.label]));
  const lines = ["graph TD"];
  for (const edge of data.edges) {
    const from = idMap.get(edge.from) || edge.from;
    const to = idMap.get(edge.to) || edge.to;
    lines.push(`    ${edge.from}["${from}"] --> ${edge.to}["${to}"]`);
  }

  const diagram = lines.join("\n");
  const { svg } = await mermaid.render("graph", diagram);
  mermaidDiv.innerHTML = svg;
}

form.onsubmit = async (e) => {
  e.preventDefault();
  const type = document.getElementById("type").value;

  try {
    if (type === "node") {
      const id = document.getElementById("id").value.trim();
      const label = document.getElementById("label").value.trim();
      if (!id || !label) throw new Error("Node ID and Label are required.");

      await fetch("/nodes", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ id, label }),
      });
    }

    if (type === "edge") {
      const from = document.getElementById("from").value.trim();
      const to = document.getElementById("to").value.trim();
      if (!from || !to) throw new Error("Both 'from' and 'to' are required.");

      await fetch("/edges", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ from, to }),
      });
    }

    messageDiv.textContent = "✅ Added successfully!";
    form.reset();
    fetchGraphAndRender();
  } catch (err) {
    messageDiv.textContent = "❌ Error: " + err.message;
  }
};

window.onload = fetchGraphAndRender;
