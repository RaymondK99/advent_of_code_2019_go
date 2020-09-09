package main

import (
	"fmt"
	"testing"
)

func Test2017_day5_part1_test1(t *testing.T)  {
	res := day5Part1_2017("0\n3\n0\n1\n-3")
	fmt.Println("res = ", res)
}

func Test2017_day5_part1(t *testing.T)  {
	input := "0\n0\n0\n1\n2\n-1\n0\n-5\n-6\n-5\n-3\n-3\n-8\n-3\n-7\n-8\n-4\n-9\n1\n-3\n-17\n0\n-18\n-14\n1\n-13\n-4\n-22\n-7\n-8\n-25\n-23\n-11\n-3\n-9\n-3\n-26\n-7\n-24\n-15\n-17\n2\n-18\n-39\n-6\n-28\n-27\n-38\n-47\n-17\n-16\n-36\n-34\n-3\n-40\n-11\n-45\n-11\n-34\n2\n-27\n-38\n-7\n-48\n-14\n-35\n-37\n-31\n-42\n-8\n0\n-39\n-20\n-50\n-9\n-7\n-18\n-6\n-65\n-44\n-68\n-67\n-61\n-19\n-3\n0\n-66\n-63\n-29\n-4\n-29\n-77\n-51\n-72\n-91\n-84\n-38\n-93\n0\n-7\n-15\n-78\n-68\n-59\n-68\n-47\n-60\n-101\n-30\n-69\n-4\n-52\n-4\n-60\n-30\n-65\n-44\n-108\n-70\n-43\n-80\n-102\n-25\n-58\n-60\n-13\n-81\n-17\n-51\n-32\n-93\n-38\n-2\n-16\n-8\n-1\n-94\n-55\n-122\n-77\n-75\n-25\n-30\n-55\n-60\n-23\n-49\n0\n-119\n-90\n-129\n-35\n-8\n-18\n-79\n-65\n-16\n-100\n-143\n-148\n-39\n-106\n-140\n-102\n-126\n-36\n-41\n-20\n-23\n-140\n-138\n-112\n-49\n-60\n-7\n-152\n-162\n-142\n-1\n-38\n-15\n-4\n-7\n-65\n0\n-114\n-150\n-160\n-162\n-29\n-164\n-32\n-60\n-60\n-97\n-130\n-97\n-172\n-25\n-185\n-35\n-77\n-138\n-188\n-182\n-137\n-122\n-133\n-6\n-37\n-199\n-108\n-53\n-188\n-60\n-30\n-160\n-160\n-23\n-176\n-10\n-23\n-6\n-116\n-216\n-155\n-107\n-66\n-55\n-107\n-132\n-91\n-152\n-127\n-142\n-90\n-54\n-40\n-86\n-151\n-123\n-120\n-37\n-84\n-86\n-205\n-117\n-179\n-32\n-26\n-228\n-232\n-116\n-67\n-122\n-125\n-61\n-2\n2\n-216\n-63\n-31\n-147\n-84\n-113\n-122\n-39\n-5\n-47\n-200\n-129\n-224\n-246\n-8\n-112\n-140\n-65\n-259\n-110\n-164\n-85\n-97\n-129\n-90\n-124\n-6\n-228\n-143\n-78\n-166\n-150\n-165\n-262\n-178\n-6\n-44\n-44\n2\n-57\n-43\n-116\n-197\n-93\n-221\n-160\n-96\n-121\n-127\n-162\n-181\n-45\n-40\n-139\n-194\n-311\n-291\n-104\n-217\n-103\n-28\n-207\n-21\n-97\n-91\n-224\n-257\n-278\n-176\n-8\n-93\n-103\n-35\n-33\n-320\n-70\n-254\n-115\n-59\n-171\n-45\n-143\n-250\n-187\n-181\n-146\n-210\n-66\n-297\n-188\n-79\n-78\n-309\n-186\n-272\n-62\n-41\n-274\n-323\n-174\n-344\n-45\n-146\n-110\n-283\n-135\n-24\n-262\n-39\n-41\n-112\n-190\n-214\n-314\n-224\n-338\n-167\n-240\n-95\n-231\n-154\n-281\n-52\n-282\n-274\n-332\n-311\n-9\n-159\n-191\n-80\n-299\n-354\n-250\n-232\n-98\n-132\n-72\n-61\n-318\n-255\n-194\n-282\n-289\n-114\n-131\n-267\n-276\n-65\n-238\n-335\n-124\n-7\n-339\n-295\n-111\n-260\n-76\n-82\n-206\n-383\n-251\n-230\n-139\n-109\n-273\n-222\n-302\n-359\n-23\n-326\n-283\n-75\n-186\n-270\n-38\n-208\n-127\n-63\n-265\n-215\n-308\n-350\n-175\n-358\n-2\n-166\n-343\n-421\n-253\n-34\n-177\n-426\n-173\n-127\n-108\n-93\n-412\n-102\n-216\n-35\n-254\n-85\n-238\n-401\n-268\n-212\n-393\n-332\n-30\n-374\n-301\n-39\n-449\n-59\n-453\n-63\n-301\n-251\n-396\n-78\n-79\n-279\n-171\n-471\n-236\n-56\n-101\n-129\n-50\n2\n-213\n-127\n-107\n-90\n-472\n-473\n-435\n-17\n-490\n-354\n-122\n-65\n-181\n-303\n-319\n-133\n-350\n-380\n-60\n-200\n-55\n-44\n-482\n-501\n-82\n-51\n-361\n-167\n-7\n-493\n-68\n-488\n-18\n-487\n-250\n-179\n-143\n-307\n-85\n-494\n-178\n-477\n-453\n-519\n-143\n-104\n-523\n-282\n-16\n-497\n-358\n-209\n-292\n-337\n-500\n-255\n-167\n-527\n-184\n-451\n-305\n-277\n-453\n-296\n-441\n-410\n-521\n-431\n-215\n-23\n-53\n-88\n-296\n-28\n-454\n-285\n-259\n-513\n-526\n-164\n-4\n-90\n-22\n-274\n-432\n-487\n-372\n-130\n-162\n-317\n-401\n-390\n-16\n-162\n-409\n-509\n-256\n-274\n-371\n-331\n-430\n-10\n-562\n-516\n-246\n-399\n-122\n-328\n-134\n-211\n-443\n-321\n-290\n-579\n-101\n-48\n-208\n-503\n-138\n-43\n-33\n-278\n-8\n-278\n-515\n-111\n-164\n-503\n-69\n-538\n-450\n-290\n-434\n-194\n-477\n-284\n-324\n-292\n-582\n-267\n-38\n-217\n-489\n-102\n-142\n-50\n-511\n-606\n-254\n-404\n-319\n-261\n-191\n-26\n-376\n-13\n-355\n-483\n-133\n-573\n-136\n-508\n-70\n-249\n-86\n-353\n-280\n-418\n-582\n-371\n-652\n-643\n-369\n-522\n-33\n-456\n-226\n-600\n-621\n-394\n-516\n-52\n-542\n-548\n-524\n-4\n-202\n-102\n-405\n-203\n-101\n-76\n-160\n-30\n-314\n-45\n-211\n-223\n-35\n-672\n-640\n-576\n-198\n-131\n-478\n-334\n-44\n-489\n-351\n-180\n-242\n-473\n-462\n-519\n-651\n-608\n-360\n-626\n-262\n-403\n-50\n-535\n-284\n-387\n-227\n-523\n-152\n-454\n-687\n-331\n-201\n-246\n-533\n-667\n-71\n-34\n-24\n-190\n-227\n-528\n-413\n-302\n-687\n-719\n-714\n-466\n-432\n-460\n-83\n-703\n-401\n-642\n-146\n-350\n-128\n-549\n-579\n-83\n-747\n-219\n-612\n-410\n-694\n-643\n-268\n-323\n-684\n-677\n-290\n-6\n-78\n-627\n-751\n-360\n-76\n-249\n-724\n-154\n-106\n-290\n-669\n-51\n-290\n-634\n-740\n-605\n-645\n-713\n-214\n-767\n-752\n-689\n-174\n-487\n-386\n-561\n-394\n-605\n-4\n-7\n-180\n-781\n-75\n-462\n-471\n-723\n-709\n-30\n-434\n-771\n-652\n-8\n-796\n-748\n-87\n-463\n-276\n-436\n-477\n-116\n-86\n-600\n-476\n-499\n-173\n-219\n-85\n-458\n-225\n-144\n-405\n-688\n-570\n-45\n-65\n-741\n-184\n-504\n-265\n-336\n-130\n-361\n-95\n-528\n-532\n-186\n-461\n-189\n-682\n-306\n-223\n-75\n-11\n-640\n-313\n-539\n-548\n-594\n-203\n-605\n-277\n-356\n-822\n-242\n-282\n-219\n-441\n-72\n-558\n-832\n-322\n-572\n-44\n-660\n-736\n-754\n-99\n-192\n-482\n-248\n-296\n-799\n-418\n-759\n-114\n-507\n-10\n-717\n-778\n-279\n-775\n-787\n-864\n-554\n-525\n-632\n-312\n-183\n-267\n-735\n-342\n-453\n-152\n-882\n-95\n-538\n-227\n-336\n-719\n-254\n-481\n-481\n-591\n-739\n-783\n-409\n-468\n-328\n-429\n-271\n-425\n-246\n-476\n-822\n-757\n-771\n-2\n-41\n-676\n-821\n-288\n-386\n-749\n-528\n-777\n-916\n-312\n-648\n-699\n-26\n-258\n-722\n-513\n-706\n-111\n-631\n-745\n-148\n-710\n-147\n-470\n-312\n-314\n-448\n-530\n-674\n-283\n-657\n-404\n-483\n-712\n-278\n-154\n-841\n-152\n-506\n-537\n-226\n-649\n-711\n-808\n-331\n-196\n-877\n-432\n-847\n-745\n-669\n-71\n-692\n-175\n-438\n-83\n-826\n-64\n-611\n-528\n-553\n-664\n-754\n-678\n-503\n-362\n-365\n-710\n-339\n-805\n-47\n-298\n-274\n-551\n-901\n-581\n-562\n-368\n-6\n-559\n-964\n-994\n-404\n-148\n-383\n-154\n-518\n-350\n-309\n-669\n-689\n-780\n-457\n-88\n-142\n-522\n-177\n-115\n-935\n-893\n-306\n-481\n-912\n-745\n-691\n-763\n-674\n-246\n-994\n-497\n-686\n-474\n-661\n-385\n-9\n-310\n-710\n-991\n-51\n-361\n-314\n-346\n-355\n-872\n-622\n-998\n-586\n-631\n-1023\n-798\n-789\n-55\n-166\n-658\n-818\n-850\n-590\n-1035\n-274\n-524\n-913\n-183\n-383\n-297\n-378\n-186\n-858\n-369\n-191\n-815\n-648\n-367\n-460\n-879\n-1052\n-476\n-382\n-455\n-892\n-57\n-54\n-523\n-824\n-1030\n-675\n-980\n-818\n-629\n-201\n-939\n-586\n-913\n-858"
	res := day5Part1_2017(input)
	fmt.Println("res = ", res)
}

func Test2017_day5_part2_test1(t *testing.T)  {
	res := day5Part2_2017("0\n3\n0\n1\n-3")
	fmt.Println("res = ", res)
}

func Test2017_day5_part2(t *testing.T)  {
	input := "0\n0\n0\n1\n2\n-1\n0\n-5\n-6\n-5\n-3\n-3\n-8\n-3\n-7\n-8\n-4\n-9\n1\n-3\n-17\n0\n-18\n-14\n1\n-13\n-4\n-22\n-7\n-8\n-25\n-23\n-11\n-3\n-9\n-3\n-26\n-7\n-24\n-15\n-17\n2\n-18\n-39\n-6\n-28\n-27\n-38\n-47\n-17\n-16\n-36\n-34\n-3\n-40\n-11\n-45\n-11\n-34\n2\n-27\n-38\n-7\n-48\n-14\n-35\n-37\n-31\n-42\n-8\n0\n-39\n-20\n-50\n-9\n-7\n-18\n-6\n-65\n-44\n-68\n-67\n-61\n-19\n-3\n0\n-66\n-63\n-29\n-4\n-29\n-77\n-51\n-72\n-91\n-84\n-38\n-93\n0\n-7\n-15\n-78\n-68\n-59\n-68\n-47\n-60\n-101\n-30\n-69\n-4\n-52\n-4\n-60\n-30\n-65\n-44\n-108\n-70\n-43\n-80\n-102\n-25\n-58\n-60\n-13\n-81\n-17\n-51\n-32\n-93\n-38\n-2\n-16\n-8\n-1\n-94\n-55\n-122\n-77\n-75\n-25\n-30\n-55\n-60\n-23\n-49\n0\n-119\n-90\n-129\n-35\n-8\n-18\n-79\n-65\n-16\n-100\n-143\n-148\n-39\n-106\n-140\n-102\n-126\n-36\n-41\n-20\n-23\n-140\n-138\n-112\n-49\n-60\n-7\n-152\n-162\n-142\n-1\n-38\n-15\n-4\n-7\n-65\n0\n-114\n-150\n-160\n-162\n-29\n-164\n-32\n-60\n-60\n-97\n-130\n-97\n-172\n-25\n-185\n-35\n-77\n-138\n-188\n-182\n-137\n-122\n-133\n-6\n-37\n-199\n-108\n-53\n-188\n-60\n-30\n-160\n-160\n-23\n-176\n-10\n-23\n-6\n-116\n-216\n-155\n-107\n-66\n-55\n-107\n-132\n-91\n-152\n-127\n-142\n-90\n-54\n-40\n-86\n-151\n-123\n-120\n-37\n-84\n-86\n-205\n-117\n-179\n-32\n-26\n-228\n-232\n-116\n-67\n-122\n-125\n-61\n-2\n2\n-216\n-63\n-31\n-147\n-84\n-113\n-122\n-39\n-5\n-47\n-200\n-129\n-224\n-246\n-8\n-112\n-140\n-65\n-259\n-110\n-164\n-85\n-97\n-129\n-90\n-124\n-6\n-228\n-143\n-78\n-166\n-150\n-165\n-262\n-178\n-6\n-44\n-44\n2\n-57\n-43\n-116\n-197\n-93\n-221\n-160\n-96\n-121\n-127\n-162\n-181\n-45\n-40\n-139\n-194\n-311\n-291\n-104\n-217\n-103\n-28\n-207\n-21\n-97\n-91\n-224\n-257\n-278\n-176\n-8\n-93\n-103\n-35\n-33\n-320\n-70\n-254\n-115\n-59\n-171\n-45\n-143\n-250\n-187\n-181\n-146\n-210\n-66\n-297\n-188\n-79\n-78\n-309\n-186\n-272\n-62\n-41\n-274\n-323\n-174\n-344\n-45\n-146\n-110\n-283\n-135\n-24\n-262\n-39\n-41\n-112\n-190\n-214\n-314\n-224\n-338\n-167\n-240\n-95\n-231\n-154\n-281\n-52\n-282\n-274\n-332\n-311\n-9\n-159\n-191\n-80\n-299\n-354\n-250\n-232\n-98\n-132\n-72\n-61\n-318\n-255\n-194\n-282\n-289\n-114\n-131\n-267\n-276\n-65\n-238\n-335\n-124\n-7\n-339\n-295\n-111\n-260\n-76\n-82\n-206\n-383\n-251\n-230\n-139\n-109\n-273\n-222\n-302\n-359\n-23\n-326\n-283\n-75\n-186\n-270\n-38\n-208\n-127\n-63\n-265\n-215\n-308\n-350\n-175\n-358\n-2\n-166\n-343\n-421\n-253\n-34\n-177\n-426\n-173\n-127\n-108\n-93\n-412\n-102\n-216\n-35\n-254\n-85\n-238\n-401\n-268\n-212\n-393\n-332\n-30\n-374\n-301\n-39\n-449\n-59\n-453\n-63\n-301\n-251\n-396\n-78\n-79\n-279\n-171\n-471\n-236\n-56\n-101\n-129\n-50\n2\n-213\n-127\n-107\n-90\n-472\n-473\n-435\n-17\n-490\n-354\n-122\n-65\n-181\n-303\n-319\n-133\n-350\n-380\n-60\n-200\n-55\n-44\n-482\n-501\n-82\n-51\n-361\n-167\n-7\n-493\n-68\n-488\n-18\n-487\n-250\n-179\n-143\n-307\n-85\n-494\n-178\n-477\n-453\n-519\n-143\n-104\n-523\n-282\n-16\n-497\n-358\n-209\n-292\n-337\n-500\n-255\n-167\n-527\n-184\n-451\n-305\n-277\n-453\n-296\n-441\n-410\n-521\n-431\n-215\n-23\n-53\n-88\n-296\n-28\n-454\n-285\n-259\n-513\n-526\n-164\n-4\n-90\n-22\n-274\n-432\n-487\n-372\n-130\n-162\n-317\n-401\n-390\n-16\n-162\n-409\n-509\n-256\n-274\n-371\n-331\n-430\n-10\n-562\n-516\n-246\n-399\n-122\n-328\n-134\n-211\n-443\n-321\n-290\n-579\n-101\n-48\n-208\n-503\n-138\n-43\n-33\n-278\n-8\n-278\n-515\n-111\n-164\n-503\n-69\n-538\n-450\n-290\n-434\n-194\n-477\n-284\n-324\n-292\n-582\n-267\n-38\n-217\n-489\n-102\n-142\n-50\n-511\n-606\n-254\n-404\n-319\n-261\n-191\n-26\n-376\n-13\n-355\n-483\n-133\n-573\n-136\n-508\n-70\n-249\n-86\n-353\n-280\n-418\n-582\n-371\n-652\n-643\n-369\n-522\n-33\n-456\n-226\n-600\n-621\n-394\n-516\n-52\n-542\n-548\n-524\n-4\n-202\n-102\n-405\n-203\n-101\n-76\n-160\n-30\n-314\n-45\n-211\n-223\n-35\n-672\n-640\n-576\n-198\n-131\n-478\n-334\n-44\n-489\n-351\n-180\n-242\n-473\n-462\n-519\n-651\n-608\n-360\n-626\n-262\n-403\n-50\n-535\n-284\n-387\n-227\n-523\n-152\n-454\n-687\n-331\n-201\n-246\n-533\n-667\n-71\n-34\n-24\n-190\n-227\n-528\n-413\n-302\n-687\n-719\n-714\n-466\n-432\n-460\n-83\n-703\n-401\n-642\n-146\n-350\n-128\n-549\n-579\n-83\n-747\n-219\n-612\n-410\n-694\n-643\n-268\n-323\n-684\n-677\n-290\n-6\n-78\n-627\n-751\n-360\n-76\n-249\n-724\n-154\n-106\n-290\n-669\n-51\n-290\n-634\n-740\n-605\n-645\n-713\n-214\n-767\n-752\n-689\n-174\n-487\n-386\n-561\n-394\n-605\n-4\n-7\n-180\n-781\n-75\n-462\n-471\n-723\n-709\n-30\n-434\n-771\n-652\n-8\n-796\n-748\n-87\n-463\n-276\n-436\n-477\n-116\n-86\n-600\n-476\n-499\n-173\n-219\n-85\n-458\n-225\n-144\n-405\n-688\n-570\n-45\n-65\n-741\n-184\n-504\n-265\n-336\n-130\n-361\n-95\n-528\n-532\n-186\n-461\n-189\n-682\n-306\n-223\n-75\n-11\n-640\n-313\n-539\n-548\n-594\n-203\n-605\n-277\n-356\n-822\n-242\n-282\n-219\n-441\n-72\n-558\n-832\n-322\n-572\n-44\n-660\n-736\n-754\n-99\n-192\n-482\n-248\n-296\n-799\n-418\n-759\n-114\n-507\n-10\n-717\n-778\n-279\n-775\n-787\n-864\n-554\n-525\n-632\n-312\n-183\n-267\n-735\n-342\n-453\n-152\n-882\n-95\n-538\n-227\n-336\n-719\n-254\n-481\n-481\n-591\n-739\n-783\n-409\n-468\n-328\n-429\n-271\n-425\n-246\n-476\n-822\n-757\n-771\n-2\n-41\n-676\n-821\n-288\n-386\n-749\n-528\n-777\n-916\n-312\n-648\n-699\n-26\n-258\n-722\n-513\n-706\n-111\n-631\n-745\n-148\n-710\n-147\n-470\n-312\n-314\n-448\n-530\n-674\n-283\n-657\n-404\n-483\n-712\n-278\n-154\n-841\n-152\n-506\n-537\n-226\n-649\n-711\n-808\n-331\n-196\n-877\n-432\n-847\n-745\n-669\n-71\n-692\n-175\n-438\n-83\n-826\n-64\n-611\n-528\n-553\n-664\n-754\n-678\n-503\n-362\n-365\n-710\n-339\n-805\n-47\n-298\n-274\n-551\n-901\n-581\n-562\n-368\n-6\n-559\n-964\n-994\n-404\n-148\n-383\n-154\n-518\n-350\n-309\n-669\n-689\n-780\n-457\n-88\n-142\n-522\n-177\n-115\n-935\n-893\n-306\n-481\n-912\n-745\n-691\n-763\n-674\n-246\n-994\n-497\n-686\n-474\n-661\n-385\n-9\n-310\n-710\n-991\n-51\n-361\n-314\n-346\n-355\n-872\n-622\n-998\n-586\n-631\n-1023\n-798\n-789\n-55\n-166\n-658\n-818\n-850\n-590\n-1035\n-274\n-524\n-913\n-183\n-383\n-297\n-378\n-186\n-858\n-369\n-191\n-815\n-648\n-367\n-460\n-879\n-1052\n-476\n-382\n-455\n-892\n-57\n-54\n-523\n-824\n-1030\n-675\n-980\n-818\n-629\n-201\n-939\n-586\n-913\n-858"
	res := day5Part2_2017(input)
	fmt.Println("res = ", res)
}
