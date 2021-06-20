data "archive_file" "mad_libs" {
  depends_on  = [local_file.mad_libs] #A
  type        = "zip"
  source_dir  = "${path.module}/madlibs"  #C
  output_path = "${path.cwd}/madlibs.zip" #B
}
