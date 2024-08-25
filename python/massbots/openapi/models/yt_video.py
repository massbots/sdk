# coding: utf-8

"""
    massbots.dl

    Download API for YouTube videos with Telegram integration. <br> - <ins>[API Support](https://t.me/MassbotsDevTeamBot)</ins> - <ins>[Terms of service](https://telegra.ph/massbotsdl-08-25)</ins> - <ins>[SDK for Go and Python](https://github.com/massbots/sdk)</ins> - <ins>[Status page](https://uptime.massbots.xyz)</ins>

    The version of the OpenAPI document: 0.3.5
    Generated by OpenAPI Generator (https://openapi-generator.tech)

    Do not edit the class manually.
"""  # noqa: E501


from __future__ import annotations
import pprint
import re  # noqa: F401
import json

from pydantic import BaseModel, ConfigDict, StrictInt, StrictStr
from typing import Any, ClassVar, Dict, List, Optional
from massbots.openapi.models.yt_thumbnail import YtThumbnail
from typing import Optional, Set
from typing_extensions import Self

class YtVideo(BaseModel):
    """
    YouTube video
    """ # noqa: E501
    id: Optional[StrictStr] = None
    title: Optional[StrictStr] = None
    description: Optional[StrictStr] = None
    url: Optional[StrictStr] = None
    published_at: Optional[StrictStr] = None
    category_id: Optional[StrictStr] = None
    channel_id: Optional[StrictStr] = None
    channel_title: Optional[StrictStr] = None
    channel_url: Optional[StrictStr] = None
    comment_count: Optional[StrictInt] = None
    like_count: Optional[StrictInt] = None
    view_count: Optional[StrictInt] = None
    thumbnails: Optional[Dict[str, YtThumbnail]] = None
    __properties: ClassVar[List[str]] = ["id", "title", "description", "url", "published_at", "category_id", "channel_id", "channel_title", "channel_url", "comment_count", "like_count", "view_count", "thumbnails"]

    model_config = ConfigDict(
        populate_by_name=True,
        validate_assignment=True,
        protected_namespaces=(),
    )


    def to_str(self) -> str:
        """Returns the string representation of the model using alias"""
        return pprint.pformat(self.model_dump(by_alias=True))

    def to_json(self) -> str:
        """Returns the JSON representation of the model using alias"""
        # TODO: pydantic v2: use .model_dump_json(by_alias=True, exclude_unset=True) instead
        return json.dumps(self.to_dict())

    @classmethod
    def from_json(cls, json_str: str) -> Optional[Self]:
        """Create an instance of YtVideo from a JSON string"""
        return cls.from_dict(json.loads(json_str))

    def to_dict(self) -> Dict[str, Any]:
        """Return the dictionary representation of the model using alias.

        This has the following differences from calling pydantic's
        `self.model_dump(by_alias=True)`:

        * `None` is only added to the output dict for nullable fields that
          were set at model initialization. Other fields with value `None`
          are ignored.
        """
        excluded_fields: Set[str] = set([
        ])

        _dict = self.model_dump(
            by_alias=True,
            exclude=excluded_fields,
            exclude_none=True,
        )
        # override the default output from pydantic by calling `to_dict()` of each value in thumbnails (dict)
        _field_dict = {}
        if self.thumbnails:
            for _key_thumbnails in self.thumbnails:
                if self.thumbnails[_key_thumbnails]:
                    _field_dict[_key_thumbnails] = self.thumbnails[_key_thumbnails].to_dict()
            _dict['thumbnails'] = _field_dict
        return _dict

    @classmethod
    def from_dict(cls, obj: Optional[Dict[str, Any]]) -> Optional[Self]:
        """Create an instance of YtVideo from a dict"""
        if obj is None:
            return None

        if not isinstance(obj, dict):
            return cls.model_validate(obj)

        _obj = cls.model_validate({
            "id": obj.get("id"),
            "title": obj.get("title"),
            "description": obj.get("description"),
            "url": obj.get("url"),
            "published_at": obj.get("published_at"),
            "category_id": obj.get("category_id"),
            "channel_id": obj.get("channel_id"),
            "channel_title": obj.get("channel_title"),
            "channel_url": obj.get("channel_url"),
            "comment_count": obj.get("comment_count"),
            "like_count": obj.get("like_count"),
            "view_count": obj.get("view_count"),
            "thumbnails": dict(
                (_k, YtThumbnail.from_dict(_v))
                for _k, _v in obj["thumbnails"].items()
            )
            if obj.get("thumbnails") is not None
            else None
        })
        return _obj


