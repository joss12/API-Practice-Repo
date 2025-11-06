import { graphStore } from "./graph/GraphStore";

graphStore.addNode({ id: "react", label: "ReactJS" });
graphStore.addNode({ id: "jsx", label: "JSX Syntax" });
graphStore.addEdge({ from: "react", to: "jsx" });

console.log(graphStore.getMap());
