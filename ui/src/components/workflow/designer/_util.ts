import { FlowNodeBaseType, type FlowNodeEntity } from "@flowgram.ai/fixed-layout-editor";

import { type WorkflowNode as _WorkflowNode, duplicateNode as _duplicateNode } from "@/domain/workflow";

import { type NodeJSON, NodeType } from "./nodes/typings";

/**
 * 克隆节点 JSON 对象。节点及其子节点 ID 均会重新分配。
 * @param {NodeJSON} node
 * @param {Object} options
 * @returns {NodeJSON}
 */
export const duplicateNodeJSON = (node: NodeJSON, options?: { withCopySuffix?: boolean }) => {
  return _duplicateNode(node as _WorkflowNode, options);
};

/**
 * 获取指定节点到根节点为止的所有前序节点。不包括自身和根节点或开始节点。
 * @param {FlowNodeEntity} node
 * @returns {FlowNodeEntity[]}
 */
export const getAllPreviousNodes = (node: FlowNodeEntity): FlowNodeEntity[] => {
  if (node == null) return [];

  // TODO: 不应该获取到旁路分支
  // // 先获取单一链路（即不包含分支）的全部节点
  // const chains: FlowNodeEntity[] = [];
  // let chain: FlowNodeEntity | undefined = node;
  // while (chain) {
  //   if (chain.isStart || chain.flowNodeType === FlowNodeBaseType.ROOT) {
  //     break;
  //   }

  //   chains.push(chain);
  //   chain = chain.pre ?? chain.parent;
  // }

  // 再获取实际的全部节点
  const visited = new Set<string>();
  const result: FlowNodeEntity[] = [];
  let current: FlowNodeEntity | undefined = node;
  while (current) {
    if (current.isStart || current.flowNodeType === FlowNodeBaseType.ROOT) {
      break;
    }

    if (current.flowNodeType === NodeType.Condition) {
      /**
       * condition
       *   blockIcon
       *   inlineBlocks
       *     branchBlock_1
       *       blockOrderIcon
       *       ...
       *     branchBlock_2
       *       blockOrderIcon
       *       ...
       */
      current.lastBlock?.blocks?.forEach((block) => {
        block.allChildren?.forEach((child) => {
          if (!visited.has(child.id)) {
            visited.add(child.id);
            result.push(child);
          }
        });
      });
    } else if (current.flowNodeType === NodeType.TryCatch) {
      /**
       * tryCatch
       *   blockIcon
       *   mainInlineBlocks
       *     tryBlock
       *       trySlot
       *       ...
       *     catchInlineBlocks
       *       catchBlock_1
       *         blockOrderIcon
       *         ...
       *         end
       *       catchBlock_2
       *         blockOrderIcon
       *         ...
       *         end
       */
      current.lastBlock?.blocks?.forEach((block) => {
        block.allChildren?.forEach((child) => {
          if (!visited.has(child.id)) {
            visited.add(child.id);
            result.push(child);
          }
        });
      });
    }

    if (!visited.has(current.id)) {
      visited.add(current.id);
      result.push(current);
    }
    current = current.pre ?? current.parent;
  }

  const prevNodes = result.filter((e) => {
    if (e.id === node.id) return false;
    if (e.isTypeOrExtendType(FlowNodeBaseType.BLOCK_ICON)) return false;
    if (e.isTypeOrExtendType(FlowNodeBaseType.BLOCK_ORDER_ICON)) return false;
    if (e.isTypeOrExtendType(FlowNodeBaseType.INLINE_BLOCKS)) return false;
    if (e.isTypeOrExtendType("trySlot")) return false;

    return true;
  });
  // console.log(node.document.toString());
  // console.log(node.document.root);
  // console.log(prevNodes);
  return prevNodes;
};
