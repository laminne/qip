import { FindUserService } from "../../../service/user/find_user_service.js";
import { AsyncResult, Failure, Success } from "../../../helpers/result.js";
import { Snowflake } from "../../../helpers/id_generator.js";

export class PersonController {
  private readonly findService: FindUserService;

  constructor(findService: FindUserService) {
    this.findService = findService;
  }

  async Handle(uHandle: string): AsyncResult<APActor, Error> {
    const user = await this.findService.FindByID(uHandle as Snowflake);
    if (user.isFailure()) {
      return new Failure(user.value);
    }
    if (!user.value.isLocalUser) {
      return new Failure(new Error("not local user"));
    }

    const res: APActor = {
      "@context": [
        "https://www.w3.org/ns/activitystreams",
        "https://w3id.org/security/v1",
      ],
      id: `https://wrt2.laminne33569.net/users/${uHandle}`,

      type: "Person",
      discoverable: true,
      preferredUsername: user.value.handle,
      name: user.value.nickName,
      inbox: user.value.apData.inboxURL,
      outbox: user.value.apData.outboxURL,
      summary: user.value.bio,
      icon: {
        name: null,
        sensitive: false,
        type: "Image",
        url: user.value.iconImageURL,
      },
      image: {
        name: null,
        sensitive: false,
        type: "Image",
        url: user.value.headerImageURL,
      },
      publicKey: {
        id: `https://wrt2.laminne33569.net/users/${uHandle}#main-key`,
        type: "Key",
        owner: `https://wrt2.laminne33569.net/users/${uHandle}`,
        publicKeyPem: user.value.apData.publicKey,
      },
    };

    return new Success(res);
  }
}

export interface APActor {
  "@context": string[];
  id: string;
  type: string;
  preferredUsername: string;
  inbox: string;
  outbox: string;
  discoverable: boolean;
  publicKey: {
    id: string;
    type: string;
    owner: string;
    publicKeyPem: string;
  };
  name: string;
  summary: string;
  icon: {
    type: string;
    url: string;
    sensitive: boolean;
    name: string | null;
  };
  image: {
    type: string;
    url: string;
    sensitive: boolean;
    name: string | null;
  };
}
