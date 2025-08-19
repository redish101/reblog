/* eslint-disable */
/* tslint:disable */
// @ts-nocheck
/*
 * ---------------------------------------------------------------
 * ## THIS FILE WAS GENERATED VIA SWAGGER-TYPESCRIPT-API        ##
 * ##                                                           ##
 * ## AUTHOR: acacode                                           ##
 * ## SOURCE: https://github.com/acacode/swagger-typescript-api ##
 * ---------------------------------------------------------------
 */

export interface CommonFailureResponse {
  error?: string;
}

export interface GormDeletedAt {
  time?: string;
  /** Valid is true if Time is not NULL */
  valid?: boolean;
}

export interface HandlerCreateOrUpdateCategoryRequest {
  description?: string;
  name?: string;
}

export interface HandlerCreateOrUpdatePostRequest {
  category?: string;
  content?: string;
  is_draft?: boolean;
  slug?: string;
  summary?: string;
  tags?: string[];
  title?: string;
}

export interface HandlerCreateOrUpdateTagRequest {
  name?: string;
}

export interface HandlerGithubCallBackResponse {
  email?: string;
  nickname?: string;
  username?: string;
}

export interface HandlerLogoutResponse {
  message?: string;
}

export interface ModelCategoryModel {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  description?: string;
  id?: number;
  name?: string;
  updated_at?: string;
}

export interface ModelPostModel {
  category?: ModelCategoryModel;
  category_id?: number;
  content?: string;
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  is_draft?: boolean;
  slug?: string;
  summary?: string;
  tags?: ModelTagModel[];
  title?: string;
  updated_at?: string;
}

export interface ModelTagModel {
  created_at?: string;
  deleted_at?: GormDeletedAt;
  id?: number;
  name?: string;
  updated_at?: string;
}

export interface StorePaginationResponseModelCategoryModel {
  /** 数据列表 */
  data?: ModelCategoryModel[];
  /** 是否有下一页 */
  has_next?: boolean;
  /** 是否有上一页 */
  has_prev?: boolean;
  /** 当前页码 */
  page?: number;
  /** 每页大小 */
  page_size?: number;
  /** 总记录数 */
  total?: number;
  /** 总页数 */
  total_pages?: number;
}

export interface StorePaginationResponseModelPostModel {
  /** 数据列表 */
  data?: ModelPostModel[];
  /** 是否有下一页 */
  has_next?: boolean;
  /** 是否有上一页 */
  has_prev?: boolean;
  /** 当前页码 */
  page?: number;
  /** 每页大小 */
  page_size?: number;
  /** 总记录数 */
  total?: number;
  /** 总页数 */
  total_pages?: number;
}

export interface StorePaginationResponseModelTagModel {
  /** 数据列表 */
  data?: ModelTagModel[];
  /** 是否有下一页 */
  has_next?: boolean;
  /** 是否有上一页 */
  has_prev?: boolean;
  /** 当前页码 */
  page?: number;
  /** 每页大小 */
  page_size?: number;
  /** 总记录数 */
  total?: number;
  /** 总页数 */
  total_pages?: number;
}

export type QueryParamsType = Record<string | number, any>;
export type ResponseFormat = keyof Omit<Body, "body" | "bodyUsed">;

export interface FullRequestParams extends Omit<RequestInit, "body"> {
  /** set parameter to `true` for call `securityWorker` for this request */
  secure?: boolean;
  /** request path */
  path: string;
  /** content type of request body */
  type?: ContentType;
  /** query params */
  query?: QueryParamsType;
  /** format of response (i.e. response.json() -> format: "json") */
  format?: ResponseFormat;
  /** request body */
  body?: unknown;
  /** base url */
  baseUrl?: string;
  /** request cancellation token */
  cancelToken?: CancelToken;
}

export type RequestParams = Omit<
  FullRequestParams,
  "body" | "method" | "query" | "path"
>;

export interface ApiConfig<SecurityDataType = unknown> {
  baseUrl?: string;
  baseApiParams?: Omit<RequestParams, "baseUrl" | "cancelToken" | "signal">;
  securityWorker?: (
    securityData: SecurityDataType | null,
  ) => Promise<RequestParams | void> | RequestParams | void;
  customFetch?: typeof fetch;
}

export interface HttpResponse<D extends unknown, E extends unknown = unknown>
  extends Response {
  data: D;
  error: E;
}

type CancelToken = Symbol | string | number;

export enum ContentType {
  Json = "application/json",
  JsonApi = "application/vnd.api+json",
  FormData = "multipart/form-data",
  UrlEncoded = "application/x-www-form-urlencoded",
  Text = "text/plain",
}

export class HttpClient<SecurityDataType = unknown> {
  public baseUrl: string = "/api/v1";
  private securityData: SecurityDataType | null = null;
  private securityWorker?: ApiConfig<SecurityDataType>["securityWorker"];
  private abortControllers = new Map<CancelToken, AbortController>();
  private customFetch = (...fetchParams: Parameters<typeof fetch>) =>
    fetch(...fetchParams);

  private baseApiParams: RequestParams = {
    credentials: "same-origin",
    headers: {},
    redirect: "follow",
    referrerPolicy: "no-referrer",
  };

  constructor(apiConfig: ApiConfig<SecurityDataType> = {}) {
    Object.assign(this, apiConfig);
  }

  public setSecurityData = (data: SecurityDataType | null) => {
    this.securityData = data;
  };

  protected encodeQueryParam(key: string, value: any) {
    const encodedKey = encodeURIComponent(key);
    return `${encodedKey}=${encodeURIComponent(typeof value === "number" ? value : `${value}`)}`;
  }

  protected addQueryParam(query: QueryParamsType, key: string) {
    return this.encodeQueryParam(key, query[key]);
  }

  protected addArrayQueryParam(query: QueryParamsType, key: string) {
    const value = query[key];
    return value.map((v: any) => this.encodeQueryParam(key, v)).join("&");
  }

  protected toQueryString(rawQuery?: QueryParamsType): string {
    const query = rawQuery || {};
    const keys = Object.keys(query).filter(
      (key) => "undefined" !== typeof query[key],
    );
    return keys
      .map((key) =>
        Array.isArray(query[key])
          ? this.addArrayQueryParam(query, key)
          : this.addQueryParam(query, key),
      )
      .join("&");
  }

  protected addQueryParams(rawQuery?: QueryParamsType): string {
    const queryString = this.toQueryString(rawQuery);
    return queryString ? `?${queryString}` : "";
  }

  private contentFormatters: Record<ContentType, (input: any) => any> = {
    [ContentType.Json]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string")
        ? JSON.stringify(input)
        : input,
    [ContentType.JsonApi]: (input: any) =>
      input !== null && (typeof input === "object" || typeof input === "string")
        ? JSON.stringify(input)
        : input,
    [ContentType.Text]: (input: any) =>
      input !== null && typeof input !== "string"
        ? JSON.stringify(input)
        : input,
    [ContentType.FormData]: (input: any) => {
      if (input instanceof FormData) {
        return input;
      }

      return Object.keys(input || {}).reduce((formData, key) => {
        const property = input[key];
        formData.append(
          key,
          property instanceof Blob
            ? property
            : typeof property === "object" && property !== null
              ? JSON.stringify(property)
              : `${property}`,
        );
        return formData;
      }, new FormData());
    },
    [ContentType.UrlEncoded]: (input: any) => this.toQueryString(input),
  };

  protected mergeRequestParams(
    params1: RequestParams,
    params2?: RequestParams,
  ): RequestParams {
    return {
      ...this.baseApiParams,
      ...params1,
      ...(params2 || {}),
      headers: {
        ...(this.baseApiParams.headers || {}),
        ...(params1.headers || {}),
        ...((params2 && params2.headers) || {}),
      },
    };
  }

  protected createAbortSignal = (
    cancelToken: CancelToken,
  ): AbortSignal | undefined => {
    if (this.abortControllers.has(cancelToken)) {
      const abortController = this.abortControllers.get(cancelToken);
      if (abortController) {
        return abortController.signal;
      }
      return void 0;
    }

    const abortController = new AbortController();
    this.abortControllers.set(cancelToken, abortController);
    return abortController.signal;
  };

  public abortRequest = (cancelToken: CancelToken) => {
    const abortController = this.abortControllers.get(cancelToken);

    if (abortController) {
      abortController.abort();
      this.abortControllers.delete(cancelToken);
    }
  };

  public request = async <T = any, E = any>({
    body,
    secure,
    path,
    type,
    query,
    format,
    baseUrl,
    cancelToken,
    ...params
  }: FullRequestParams): Promise<HttpResponse<T, E>> => {
    const secureParams =
      ((typeof secure === "boolean" ? secure : this.baseApiParams.secure) &&
        this.securityWorker &&
        (await this.securityWorker(this.securityData))) ||
      {};
    const requestParams = this.mergeRequestParams(params, secureParams);
    const queryString = query && this.toQueryString(query);
    const payloadFormatter = this.contentFormatters[type || ContentType.Json];
    const responseFormat = format || requestParams.format;

    return this.customFetch(
      `${baseUrl || this.baseUrl || ""}${path}${queryString ? `?${queryString}` : ""}`,
      {
        ...requestParams,
        headers: {
          ...(requestParams.headers || {}),
          ...(type && type !== ContentType.FormData
            ? { "Content-Type": type }
            : {}),
        },
        signal:
          (cancelToken
            ? this.createAbortSignal(cancelToken)
            : requestParams.signal) || null,
        body:
          typeof body === "undefined" || body === null
            ? null
            : payloadFormatter(body),
      },
    ).then(async (response) => {
      const r = response.clone() as HttpResponse<T, E>;
      r.data = null as unknown as T;
      r.error = null as unknown as E;

      const data = !responseFormat
        ? r
        : await response[responseFormat]()
            .then((data) => {
              if (r.ok) {
                r.data = data;
              } else {
                r.error = data;
              }
              return r;
            })
            .catch((e) => {
              r.error = e;
              return r;
            });

      if (cancelToken) {
        this.abortControllers.delete(cancelToken);
      }

      if (!response.ok) throw data;
      return data;
    });
  };
}

/**
 * @title reblog api
 * @version 1.0
 * @baseUrl /api/v1
 * @contact
 *
 * reblog 后端 api
 */
export class Api<
  SecurityDataType extends unknown,
> extends HttpClient<SecurityDataType> {
  auth = {
    /**
     * @description 跳转到 GitHub OAuth 授权页面
     *
     * @tags auth
     * @name GithubList
     * @summary GitHub OAuth 登录
     * @request GET:/auth/github
     */
    githubList: (
      query?: {
        /** 登录成功后的重定向地址 */
        redirect_uri?: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<any, string>({
        path: `/auth/github`,
        method: "GET",
        query: query,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description 处理 GitHub OAuth 回调，获取用户信息并生成 JWT token
     *
     * @tags auth
     * @name GithubCallbackList
     * @summary GitHub OAuth 回调
     * @request GET:/auth/github/callback
     */
    githubCallbackList: (
      query: {
        /** GitHub OAuth 授权码 */
        code: string;
        /** 状态参数 */
        state?: string;
      },
      params: RequestParams = {},
    ) =>
      this.request<HandlerGithubCallBackResponse, CommonFailureResponse>({
        path: `/auth/github/callback`,
        method: "GET",
        query: query,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 清除登录状态
     *
     * @tags auth
     * @name LogoutCreate
     * @summary 登出
     * @request POST:/auth/logout
     * @secure
     */
    logoutCreate: (params: RequestParams = {}) =>
      this.request<HandlerLogoutResponse, CommonFailureResponse>({
        path: `/auth/logout`,
        method: "POST",
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  categories = {
    /**
     * @description 获取分类列表
     *
     * @tags category
     * @name CategoriesList
     * @summary 获取分类列表
     * @request GET:/categories
     */
    categoriesList: (
      query?: {
        /** 页码，默认为 1 */
        page?: number;
        /** 每页数量，默认为 10 */
        page_size?: number;
      },
      params: RequestParams = {},
    ) =>
      this.request<
        StorePaginationResponseModelCategoryModel,
        CommonFailureResponse
      >({
        path: `/categories`,
        method: "GET",
        query: query,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 创建新的分类
     *
     * @tags category
     * @name CategoriesCreate
     * @summary 创建分类
     * @request POST:/categories
     * @secure
     */
    categoriesCreate: (
      body: HandlerCreateOrUpdateCategoryRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelCategoryModel, CommonFailureResponse>({
        path: `/categories`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 删除指定名称的分类
     *
     * @tags category
     * @name CategoriesDelete
     * @summary 删除分类
     * @request DELETE:/categories/{name}
     * @secure
     */
    categoriesDelete: (name: string, params: RequestParams = {}) =>
      this.request<void, CommonFailureResponse>({
        path: `/categories/${name}`,
        method: "DELETE",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description 根据分类 名称 获取分类详情
     *
     * @tags category
     * @name CategoriesDetail
     * @summary 获取分类详情
     * @request GET:/categories/{name}
     */
    categoriesDetail: (name: string, params: RequestParams = {}) =>
      this.request<ModelCategoryModel, CommonFailureResponse>({
        path: `/categories/${name}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 更新分类信息
     *
     * @tags category
     * @name CategoriesUpdate
     * @summary 更新分类
     * @request PUT:/categories/{name}
     * @secure
     */
    categoriesUpdate: (
      name: string,
      body: HandlerCreateOrUpdateCategoryRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelCategoryModel, CommonFailureResponse>({
        path: `/categories/${name}`,
        method: "PUT",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  healthz = {
    /**
     * @description 健康检查
     *
     * @tags healthz
     * @name HealthzList
     * @summary 健康检查
     * @request GET:/healthz
     */
    healthzList: (params: RequestParams = {}) =>
      this.request<boolean, any>({
        path: `/healthz`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  posts = {
    /**
     * @description 获取文章列表
     *
     * @tags post
     * @name PostsList
     * @summary 获取文章列表
     * @request GET:/posts
     */
    postsList: (
      query?: {
        /** 页码，默认为 1 */
        page?: number;
        /** 每页数量，默认为 10 */
        page_size?: number;
        /** 分类名称过滤 */
        categories?: string[];
        /** 标签名称过滤 */
        tags?: string[];
      },
      params: RequestParams = {},
    ) =>
      this.request<
        StorePaginationResponseModelPostModel,
        CommonFailureResponse
      >({
        path: `/posts`,
        method: "GET",
        query: query,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 创建文章
     *
     * @tags post
     * @name PostsCreate
     * @summary 创建文章
     * @request POST:/posts
     * @secure
     */
    postsCreate: (
      body: HandlerCreateOrUpdatePostRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelPostModel, CommonFailureResponse>({
        path: `/posts`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 删除文章
     *
     * @tags post
     * @name PostsDelete
     * @summary 删除文章
     * @request DELETE:/posts/{slug}
     * @secure
     */
    postsDelete: (slug: string, params: RequestParams = {}) =>
      this.request<void, CommonFailureResponse>({
        path: `/posts/${slug}`,
        method: "DELETE",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description 获取文章详情
     *
     * @tags post
     * @name PostsDetail
     * @summary 获取文章详情
     * @request GET:/posts/{slug}
     */
    postsDetail: (slug: string, params: RequestParams = {}) =>
      this.request<ModelPostModel, CommonFailureResponse>({
        path: `/posts/${slug}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 更新文章
     *
     * @tags post
     * @name PostsUpdate
     * @summary 更新文章
     * @request PUT:/posts/{slug}
     * @secure
     */
    postsUpdate: (
      slug: string,
      body: HandlerCreateOrUpdatePostRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelPostModel, CommonFailureResponse>({
        path: `/posts/${slug}`,
        method: "PUT",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
  tags = {
    /**
     * @description 获取标签的列表
     *
     * @tags tag
     * @name TagsList
     * @summary 获取标签列表
     * @request GET:/tags
     */
    tagsList: (
      query?: {
        /** 页码，默认为 1 */
        page?: number;
        /** 每页数量，默认为 10 */
        page_size?: number;
      },
      params: RequestParams = {},
    ) =>
      this.request<StorePaginationResponseModelTagModel, CommonFailureResponse>(
        {
          path: `/tags`,
          method: "GET",
          query: query,
          type: ContentType.Json,
          format: "json",
          ...params,
        },
      ),

    /**
     * @description 创建一个新的标签
     *
     * @tags tag
     * @name TagsCreate
     * @summary 创建标签
     * @request POST:/tags
     * @secure
     */
    tagsCreate: (
      body: HandlerCreateOrUpdateTagRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelTagModel, CommonFailureResponse>({
        path: `/tags`,
        method: "POST",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 删除指定名称的标签
     *
     * @tags tag
     * @name TagsDelete
     * @summary 删除标签
     * @request DELETE:/tags/{name}
     * @secure
     */
    tagsDelete: (name: string, params: RequestParams = {}) =>
      this.request<void, CommonFailureResponse>({
        path: `/tags/${name}`,
        method: "DELETE",
        secure: true,
        type: ContentType.Json,
        ...params,
      }),

    /**
     * @description 获取指定名称的标签详情
     *
     * @tags tag
     * @name TagsDetail
     * @summary 获取标签详情
     * @request GET:/tags/{name}
     */
    tagsDetail: (name: string, params: RequestParams = {}) =>
      this.request<ModelTagModel, CommonFailureResponse>({
        path: `/tags/${name}`,
        method: "GET",
        type: ContentType.Json,
        format: "json",
        ...params,
      }),

    /**
     * @description 更新现有标签
     *
     * @tags tag
     * @name TagsUpdate
     * @summary 更新标签
     * @request PUT:/tags/{name}
     * @secure
     */
    tagsUpdate: (
      name: string,
      body: HandlerCreateOrUpdateTagRequest,
      params: RequestParams = {},
    ) =>
      this.request<ModelTagModel, CommonFailureResponse>({
        path: `/tags/${name}`,
        method: "PUT",
        body: body,
        secure: true,
        type: ContentType.Json,
        format: "json",
        ...params,
      }),
  };
}
