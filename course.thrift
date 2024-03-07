namespace go course

struct TreeNode {
    1: string name
    2: map<string, string> data
    3: string path
    4: list<TreeNode> children
    5: TreeNode parent
}

service course {
   TreeNode getCourseProgress() (api.get = "progress")
   i64 resetCourseProgress(1:TreeNode root) (api.delete = "progress")
   i64 updateCourseProgress(1: TreeNode root) (api.put = "progress")
   
   TreeNode getCoursePlan(1: string courseId) (api.get = "plan")
   list<i64> getCoursePlanIds() (api.get = "plan/ids")
   i64 createCoursePlan() (api.post = "plan")
   i64 updateCoursePlan(2: TreeNode root) (api.put = "plan")
   i64 deleteCoursePlan(1: i64 coursePlanId) (api.delete = "plan")
}