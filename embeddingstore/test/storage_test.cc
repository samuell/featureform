/* This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/. */

#include <iostream>
#include <memory>
#include <vector>

#include <gtest/gtest.h>
#include "embeddingstore/storage.h"

using featureform::embedding::EmbeddingStorage;

namespace {

TEST(EmbeddingStorage, TestSetGet) {
  auto storage = EmbeddingStorage::New("test.db", 3);
  auto a_key = "a";
  auto a_vec = std::vector<float>{0, 1, 0};
  storage->set(a_key, a_vec);
  ASSERT_EQ(storage->get(a_key), a_vec);
}

}
