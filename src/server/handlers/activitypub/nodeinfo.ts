import { FastifyHandlerMethod } from "../../../helpers/fastify.js";
import { NodeInfoController } from "../../controller/activitypub/nodeinfo.js";

export class NodeInfoHandlers {
  private readonly controller: NodeInfoController;
  constructor(c: NodeInfoController) {
    this.controller = c;
  }

  public Handle: FastifyHandlerMethod<{ Params: object }> = async (q, r) => {
    const res = this.controller.Handle();
    return r.code(200).send(res);
  };
}

// /nodeinfo/2.0
export interface NodeInfo20 {
  version: string;
  software: {
    name: string;
    version: string;
  };
  protocols: string[];
  services: {
    inbound: string[];
    outbound: string[];
  };
  openRegistrations: boolean;
  usage: {
    users: {
      total: number;
    };
    localPosts: number;
  };
  metadata: {
    nodeName: string;
    nodeDescription: string;
    maintainer: {
      name: string;
      email: string;
    };
  };
}
