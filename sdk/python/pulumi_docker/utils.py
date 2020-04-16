# Copyright 2016-2020, Pulumi Corporation.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
import typing


def get_image_name_and_tag(base_image_name: str) -> typing.Tuple[str, typing.Optional[str]]:
    # From https://docs.docker.com/engine/reference/commandline/tag
    #
    # "A tag name must be valid ASCII and may contain lowercase and uppercase letters, digits,
    # underscores, periods and dashes. A tag name may not start with a period or a dash and may
    # contain a maximum of 128 characters."
    #
    # So it is safe for us to just look for the colon, and consume whatever follows as the tag
    # for the image.

    try:
        last_colon = base_image_name.index(":")
        return base_image_name[0: last_colon], base_image_name[last_colon + 1:]
    except ValueError:
        return base_image_name, None
