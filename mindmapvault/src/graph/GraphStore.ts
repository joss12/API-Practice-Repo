import { NodeType } from "./NodeType";
import { EdgeType } from "./EdgeType";

export class GraphStore {
  private nodes: Map<string, NodeType> = new Map();
  private edges: EdgeType[] = [];

  addNode(node: NodeType) {
    if (this.nodes.has(node.id)) {
      throw new Error(`Node '${node.id}' already exists.`);
    }
    this.nodes.set(node.id, node);
  }

  addEdge(edge: EdgeType) {
    if (!this.nodes.has(edge.from) || !this.nodes.has(edge.to)) {
      throw new Error("Both nodes must exist before linking.");
    }
    this.edges.push(edge);
  }

  getMap() {
    return {
      nodes: Array.from(this.nodes.values()),
      edges: this.edges,
    };
  }
}

export const graphStore = new GraphStore();
